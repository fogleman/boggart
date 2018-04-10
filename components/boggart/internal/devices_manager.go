package internal

import (
	"context"
	"sync"
	"time"

	"github.com/kihamo/boggart/components/boggart"
	w "github.com/kihamo/go-workers"
	"github.com/kihamo/go-workers/manager"
	"github.com/kihamo/shadow/components/workers"
	"github.com/kihamo/snitch"
	"github.com/pborman/uuid"
)

const (
	DefaultTimeoutChecker = time.Second * 2
)

type DevicesManager struct {
	mutex sync.RWMutex

	storage        *sync.Map
	workers        workers.Component
	listeners      *manager.ListenersManager
	chanChecker    chan string
	tickerChecker  *w.Ticker
	timeoutChecker time.Duration
}

func NewDevicesManager(workers workers.Component) *DevicesManager {
	m := &DevicesManager{
		storage:        new(sync.Map),
		workers:        workers,
		listeners:      manager.NewListenersManager(),
		chanChecker:    make(chan string),
		tickerChecker:  w.NewTicker(time.Minute),
		timeoutChecker: DefaultTimeoutChecker,
	}

	go m.doCheck()

	return m
}

func (m *DevicesManager) Register(device boggart.Device) string {
	id := uuid.New()
	m.RegisterWithID(id, device)
	return id
}

func (m *DevicesManager) RegisterWithID(id string, device boggart.Device) {
	m.storage.Store(id, device)
	m.listeners.AsyncTrigger(boggart.DeviceEventDeviceRegister, device, id)

	tasks := device.Tasks()
	if len(tasks) > 0 {
		for _, task := range tasks {
			m.workers.AddTask(task)
		}
	}

	listeners := device.Listeners()
	if len(listeners) > 0 {
		for _, listener := range listeners {
			m.listeners.AddListener(listener)
		}
	}

	events := device.TriggerEventChannel()
	if events != nil {
		go m.doDeviceEvents(events)
	}

	m.CheckByKeys(id)
}

func (m *DevicesManager) Device(id string) boggart.Device {
	if d, ok := m.storage.Load(id); ok {
		return d.(boggart.Device)
	}

	return nil
}

func (m *DevicesManager) Devices() map[string]boggart.Device {
	devices := make(map[string]boggart.Device, 0)

	m.storage.Range(func(key interface{}, device interface{}) bool {
		devices[key.(string)] = device.(boggart.Device)
		return true
	})

	return devices
}

func (m *DevicesManager) DevicesByTypes(types []boggart.DeviceType) map[string]boggart.Device {
	if len(types) == 0 {
		return m.Devices()
	}

	devices := make(map[string]boggart.Device, 0)

	m.storage.Range(func(key interface{}, device interface{}) bool {
		d := device.(boggart.Device)

		for _, t1 := range d.Types() {
			for _, t2 := range types {
				if t1.String() == t2.String() {
					devices[key.(string)] = d
					return true
				}
			}
		}

		return true
	})

	return devices
}

func (m *DevicesManager) Describe(ch chan<- *snitch.Description) {
	m.storage.Range(func(_ interface{}, device interface{}) bool {
		if !device.(boggart.Device).IsEnabled() {
			return true
		}

		if collector, ok := device.(snitch.Collector); ok {
			collector.Describe(ch)
		}

		return true
	})
}

func (m *DevicesManager) Collect(ch chan<- snitch.Metric) {
	m.storage.Range(func(_ interface{}, device interface{}) bool {
		if !device.(boggart.Device).IsEnabled() {
			return true
		}

		if collector, ok := device.(snitch.Collector); ok {
			collector.Collect(ch)
		}

		return true
	})
}

func (m *DevicesManager) Attach(event w.Event, listener w.Listener) {
	m.listeners.Attach(event, listener)
}

func (m *DevicesManager) DeAttach(event w.Event, listener w.Listener) {
	m.listeners.DeAttach(event, listener)
}

func (m *DevicesManager) Ready() {
	m.listeners.AsyncTrigger(boggart.DeviceEventDevicesManagerReady)
}

func (m *DevicesManager) ListenersManager() *manager.ListenersManager {
	return m.listeners
}

func (m *DevicesManager) Listeners() []w.Listener {
	return m.listeners.Listeners()
}

func (m *DevicesManager) AddListener(listener w.ListenerWithEvents) {
	m.listeners.AddListener(listener)
}

func (m *DevicesManager) GetListenerMetadata(id string) w.Metadata {
	if item := m.listeners.GetById(id); item != nil {
		return item.Metadata()
	}

	return nil
}

func (m *DevicesManager) SetCheckerTickerDuration(duration time.Duration) {
	m.tickerChecker.SetDuration(duration)
}

func (m *DevicesManager) SetCheckerTimeout(duration time.Duration) {
	m.mutex.Lock()
	m.timeoutChecker = duration
	m.mutex.Unlock()
}

func (m *DevicesManager) Check() {
	keys := make([]string, 0, 0)

	m.storage.Range(func(key interface{}, device interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})

	m.CheckByKeys(keys...)
}

func (m *DevicesManager) CheckByKeys(keys ...string) {
	go func() {
		for _, key := range keys {
			m.chanChecker <- key
		}
	}()
}

func (m *DevicesManager) doCheck() {
	for {
		select {
		case key := <-m.chanChecker:
			if device := m.Device(key); device != nil {
				m.checker(key, device)
			}

		case <-m.tickerChecker.C():
			m.Check()
		}
	}
}

func (m *DevicesManager) doDeviceEvents(ch <-chan boggart.DeviceTriggerEvent) {
	for event := range ch {
		m.listeners.AsyncTrigger(event.Event(), event.Arguments()...)
	}
}

func (m *DevicesManager) checker(key string, device boggart.Device) {
	m.mutex.RLock()
	timeout := m.timeoutChecker
	m.mutex.RUnlock()

	ctx, ctxCancel := context.WithTimeout(context.Background(), timeout)
	defer ctxCancel()

	done := make(chan bool, 1)

	go func() {
		done <- device.Ping(ctx)
	}()

	select {
	case <-ctx.Done():
		if !device.IsEnabled() {
			return
		}

		if ctx.Err() != nil && ctx.Err() != context.Canceled {
			device.Disable()
			m.listeners.AsyncTrigger(boggart.DeviceEventDeviceDisabledAfterCheck, device, key, ctx.Err())
		}

		return

	case result := <-done:
		if result == device.IsEnabled() {
			return
		}

		if !result {
			device.Disable()
			m.listeners.AsyncTrigger(boggart.DeviceEventDeviceDisabledAfterCheck, device, key, nil)
		} else {
			device.Enable()
			m.listeners.AsyncTrigger(boggart.DeviceEventDeviceEnabledAfterCheck, device, key)
		}

		return
	}
}
