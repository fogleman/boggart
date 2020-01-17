package manager

import (
	"context"
	"errors"
	"io"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/di"
	"github.com/kihamo/boggart/components/mqtt"
	w "github.com/kihamo/go-workers"
	"github.com/kihamo/go-workers/manager"
	"github.com/kihamo/go-workers/task"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/i18n"
	"github.com/kihamo/shadow/components/logging"
	"github.com/kihamo/shadow/components/workers"
	"github.com/kihamo/snitch"
	"github.com/pborman/uuid"
	"go.uber.org/multierr"
)

const (
	managerStatusNotReady = int64(iota)
	managerStatusReady
	managerStatusClose
)

type bindTask interface {
	w.Task
	SetName(string)
}

type Manager struct {
	status          int64
	storage         sync.Map
	dashboard       dashboard.Component
	i18n            i18n.Component
	mqtt            mqtt.Component
	workers         workers.Component
	logger          logging.Logger
	listeners       *manager.ListenersManager
	topicBindStatus mqtt.Topic
}

func NewManager(dashboard dashboard.Component,
	i18n i18n.Component,
	mqtt mqtt.Component,
	workers workers.Component,
	logger logging.Logger,
	listeners *manager.ListenersManager,
	topicBindStatus mqtt.Topic) *Manager {
	return &Manager{
		status:          managerStatusNotReady,
		dashboard:       dashboard,
		i18n:            i18n,
		mqtt:            mqtt,
		workers:         workers,
		logger:          logger,
		listeners:       listeners,
		topicBindStatus: topicBindStatus,
	}
}

func (m *Manager) Register(id string, bind boggart.Bind, t string, description string, tags []string, config interface{}) (boggart.BindItem, error) {
	if id == "" {
		id = uuid.New()
	} else {
		if existsBind := m.Bind(id); existsBind != nil {
			return nil, errors.New("bind item with id " + id + " already exist")
		}
	}

	bindType, err := boggart.GetBindType(t)
	if err != nil {
		return nil, err
	}

	// register widget
	if widget, ok := bindType.(boggart.BindTypeHasWidgetAssetFS); ok {
		if fs := widget.WidgetAssetFS(); fs != nil {
			name := boggart.ComponentName + "-bind-" + t

			// templates
			render := m.dashboard.Renderer()
			if !render.IsRegisterNamespace(name) {
				if err := render.RegisterNamespace(name, fs); err != nil {
					return nil, err
				}
			}

			// asset fs
			m.dashboard.RegisterAssetFS(name, fs)

			// i18n
			if m.i18n != nil {
				fs.Prefix = "locales"
				m.i18n.LoadLocaleFromFiles(name, i18n.FromAssetFS(fs))
			}
		}
	}

	bindItem := &BindItem{
		bind:        bind,
		bindType:    bindType,
		id:          id,
		t:           t,
		description: description,
		tags:        tags,
		config:      config,
	}

	m.itemStatusUpdate(bindItem, boggart.BindStatusInitializing)

	if bindSupport, ok := bind.(di.MetaContainerSupport); ok {
		bindSupport.SetMeta(di.NewMetaContainer(bindItem))
	}

	// init logger
	if bindSupport, ok := bind.(di.LoggerContainerSupport); ok {
		bindSupport.SetLogger(di.NewLoggerContainer(logging.NewLazyLogger(m.logger, m.logger.Name()+"."+id)))
	}

	// register mqtt
	if bindSupport, ok := bind.(di.MQTTContainerSupport); ok {
		bindSupport.SetMQTT(di.NewMQTTContainer(bindItem, m.mqtt))
	}

	if runner, ok := bind.(boggart.BindRunner); ok {
		if err := runner.Run(); err != nil {
			return nil, err
		}
	}

	if bindSupport, ok := bind.(di.MQTTContainerSupport); ok {
		// TODO: обвешать подписки враппером, что бы только в online можно было посылать
		for _, subscriber := range bindSupport.MQTT().Subscribers() {
			if err := m.mqtt.SubscribeSubscriber(subscriber); err != nil {
				m.itemStatusUpdate(bindItem, boggart.BindStatusUninitialized)
				return nil, err
			}
		}
	}

	tasks := make([]w.Task, 0)

	// probes
	if probe, ok := bind.(boggart.BindHasReadinessProbe); ok {
		probeTask := task.NewFunctionTask(func(ctx context.Context) (_ interface{}, err error) {
			c := make(chan error, 1)
			go func() {
				c <- probe.ReadinessProbe(ctx)
			}()

			select {
			case <-ctx.Done():
				err = ctx.Err()
			case err = <-c:
			}

			if err != nil {
				m.itemStatusUpdate(bindItem, boggart.BindStatusOffline)

				m.logger.Error("Readiness probe failure",
					"type", bindItem.Type(),
					"id", bindItem.ID(),
					"error", err.Error(),
				)
			} else {
				m.itemStatusUpdate(bindItem, boggart.BindStatusOnline)
			}

			return nil, err
		})

		probePeriod := boggart.ReadinessProbeDefaultPeriod
		probeTimeout := boggart.ReadinessProbeDefaultTimeout
		if probeConfig, ok := config.(boggart.BindConfigReadinessProbe); ok {
			probePeriod = probeConfig.ReadinessProbePeriod()
			probeTimeout = probeConfig.ReadinessProbeTimeout()
		}
		probeTask.SetRepeatInterval(probePeriod)
		probeTask.SetTimeout(probeTimeout)

		probeTask.SetRepeats(-1)
		probeTask.SetName("readiness-probe")

		bindItem.probeReadiness = probeTask
		tasks = append(tasks, bindItem.probeReadiness)
	} else {
		m.itemStatusUpdate(bindItem, boggart.BindStatusOnline)
	}

	if probe, ok := bind.(boggart.BindHasLivenessProbe); ok {
		probeTask := task.NewFunctionTask(func(ctx context.Context) (_ interface{}, err error) {
			currentStatus := bindItem.Status()
			if currentStatus != boggart.BindStatusOnline && currentStatus != boggart.BindStatusOffline {
				return nil, nil
			}

			c := make(chan error, 1)
			go func() {
				c <- probe.LivenessProbe(ctx)
			}()

			select {
			case <-ctx.Done():
				err = ctx.Err()
			case err = <-c:
			}

			if err == nil {
				return nil, nil
			}

			m.logger.Error("Liveness probe failure",
				"type", bindItem.Type(),
				"id", bindItem.ID(),
				"error", err.Error(),
			)

			if e := m.Unregister(bindItem.ID()); e != nil {
				m.logger.Error("Unregister after liveness probe failure",
					"type", bindItem.Type(),
					"id", bindItem.ID(),
					"error", e.Error(),
				)

				err = multierr.Append(err, e)
				return nil, err
			}

			if _, e := m.Register(id, bind, t, description, tags, config); e != nil {
				m.logger.Error("Register after liveness probe failure",
					"type", bindItem.Type(),
					"id", bindItem.ID(),
					"error", e.Error(),
				)

				err = multierr.Append(err, e)
				return nil, err
			}

			return nil, err
		})

		probePeriod := boggart.LivenessProbeDefaultPeriod
		probeTimeout := boggart.LivenessProbeDefaultTimeout
		if probeConfig, ok := config.(boggart.BindConfigLivenessProbe); ok {
			probePeriod = probeConfig.LivenessProbePeriod()
			probeTimeout = probeConfig.LivenessProbeTimeout()
		}
		probeTask.SetRepeatInterval(probePeriod)
		probeTask.SetTimeout(probeTimeout)

		probeTask.SetRepeats(-1)
		probeTask.SetName("liveness-probe")

		bindItem.probeLiveness = probeTask
		tasks = append(tasks, bindItem.probeLiveness)
	}

	if bindSupport, ok := bindItem.Bind().(di.WorkersContainerSupport); ok {
		bindSupport.SetWorkers(di.NewWorkersContainer(bindItem))

		tasks = append(tasks, bindSupport.Workers().Tasks()...)

		// register listeners
		for _, listener := range bindSupport.Workers().Listeners() {
			m.listeners.AddListener(listener)
		}
	}

	// register tasks
	for _, tsk := range tasks {
		if tsk, ok := tsk.(bindTask); ok {
			tsk.SetName("bind-" + id + "-" + t + "-" + tsk.Name())
		}

		m.workers.AddTask(tsk)
	}

	m.storage.Store(id, bindItem)

	m.logger.Debug("Register bind",
		"type", bindItem.Type(),
		"id", bindItem.ID(),
	)

	return bindItem, nil
}

func (m *Manager) Unregister(id string) error {
	d, ok := m.storage.Load(id)
	if !ok {
		return nil
	}

	bindItem := d.(*BindItem)

	m.itemStatusUpdate(bindItem, boggart.BindStatusRemoving)

	// unregister mqtt
	if bindSupport, ok := bindItem.Bind().(di.MQTTContainerSupport); ok {
		if err := m.mqtt.UnsubscribeSubscribers(bindSupport.MQTT().Subscribers()); err != nil {
			m.logger.Error("Unregister bind failed because unsubscribe MQTT failed",
				"type", bindItem.Type(),
				"id", bindItem.ID(),
				"error", err.Error(),
			)
		}

		bindSupport.MQTT().SetClient(nil)
	}

	// remove probes
	if bindItem.probeReadiness != nil {
		m.workers.RemoveTask(bindItem.probeReadiness)
	}
	if bindItem.probeLiveness != nil {
		m.workers.RemoveTask(bindItem.probeLiveness)
	}

	// workers
	if bindSupport, ok := bindItem.Bind().(di.WorkersContainerSupport); ok {
		// remove tasks
		for _, tsk := range bindSupport.Workers().Tasks() {
			m.workers.RemoveTask(tsk)
		}

		// remove listeners
		for _, listener := range bindSupport.Workers().Listeners() {
			m.listeners.RemoveListener(listener)
		}
	}

	m.storage.Delete(id)

	m.logger.Debug("Unregister bind",
		"type", bindItem.Type(),
		"id", bindItem.ID(),
	)

	if closer, ok := bindItem.Bind().(io.Closer); ok {
		if err := closer.Close(); err != nil {
			m.logger.Debug("Unregister bind failed because close failed",
				"type", bindItem.Type(),
				"id", bindItem.ID(),
				"error", err.Error(),
			)
		}
	}

	m.itemStatusUpdate(bindItem, boggart.BindStatusRemoved)
	return nil
}

func (m *Manager) UnregisterAll() error {
	for _, item := range m.BindItems() {
		if err := m.Unregister(item.ID()); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) Bind(id string) boggart.BindItem {
	if d, ok := m.storage.Load(id); ok {
		return d.(boggart.BindItem)
	}

	return nil
}

func (m *Manager) BindItems() BindItemsList {
	items := make([]boggart.BindItem, 0)

	m.storage.Range(func(_ interface{}, item interface{}) bool {
		items = append(items, item.(boggart.BindItem))
		return true
	})

	sort.Slice(items, func(i, j int) bool {
		if items[i].Type() == items[j].Type() {
			return items[i].ID() < items[j].ID()
		}

		return items[i].Type() < items[j].Type()
	})

	return items
}

func (m *Manager) Describe(ch chan<- *snitch.Description) {
	m.storage.Range(func(_ interface{}, item interface{}) bool {
		if collector, ok := item.(*BindItem).Bind().(snitch.Collector); ok {
			collector.Describe(ch)
		}

		return true
	})
}

func (m *Manager) Collect(ch chan<- snitch.Metric) {
	m.storage.Range(func(_ interface{}, item interface{}) bool {
		if collector, ok := item.(*BindItem).Bind().(snitch.Collector); ok {
			collector.Collect(ch)
		}

		return true
	})
}

func (m *Manager) Ready() {
	atomic.SwapInt64(&m.status, managerStatusReady)
}

func (m *Manager) Close() error {
	old := atomic.SwapInt64(&m.status, managerStatusClose)
	if old != managerStatusClose {
		return m.UnregisterAll()
	}

	return nil
}

func (m *Manager) MQTTOnConnectHandler(client mqtt.Component, restore bool) {
	for _, item := range m.BindItems() {
		topic := m.topicBindStatus.Format(item.ID())
		err := client.PublishWithoutCache(context.Background(), topic, 1, true, strings.ToLower(item.Status().String()))

		if err != nil {
			m.logger.Error("Restore publish to MQTT failed", "topic", topic, "error", err.Error())
		}
	}
}

func (m *Manager) ReadinessProbeCheck(ctx context.Context, id string) (err error) {
	if d, ok := m.storage.Load(id); ok {
		if probe := d.(*BindItem).probeReadiness; probe != nil {
			if timeout := probe.Timeout(); timeout > 0 {
				ctx, _ = context.WithTimeout(ctx, timeout)
			}

			_, err = probe.Run(ctx)
		}
	}

	return err
}

func (m *Manager) LivenessProbeCheck(ctx context.Context, id string) (err error) {
	if d, ok := m.storage.Load(id); ok {
		if probe := d.(*BindItem).probeLiveness; probe != nil {
			if timeout := probe.Timeout(); timeout > 0 {
				ctx, _ = context.WithTimeout(ctx, timeout)
			}

			_, err = probe.Run(ctx)
		}
	}

	return err
}

func (m *Manager) mqttPublish(topic mqtt.Topic, payload interface{}) {
	// при закрытии шлем синхронно, что бы блочить операцию Close компонента
	if atomic.LoadInt64(&m.status) == managerStatusClose {
		if err := m.mqtt.PublishWithoutCache(context.Background(), topic, 1, true, payload); err != nil {
			m.logger.Error("Publish to MQTT failed", "topic", topic, "error", err.Error())
		}
	} else {
		m.mqtt.PublishAsync(context.Background(), topic, 1, true, payload)
	}
}

func (m *Manager) itemStatusUpdate(item *BindItem, status boggart.BindStatus) {
	if ok := item.updateStatus(status); ok {
		m.mqttPublish(m.topicBindStatus.Format(item.id), strings.ToLower(status.String()))
	}
}
