package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"sort"
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/di"
	"github.com/kihamo/boggart/components/boggart/tasks"
	"github.com/kihamo/boggart/components/mqtt"
	"github.com/kihamo/shadow/components/dashboard"
	"gopkg.in/yaml.v2"
)

type BindYAML struct {
	Type        string
	ID          string
	Description string
	Tags        []string
	Config      interface{}
}

type BindHandler struct {
	dashboard.Handler

	componentBoggart boggart.Component
	componentMQTT    mqtt.Component
}

func NewBindHandler(b boggart.Component, m mqtt.Component) *BindHandler {
	return &BindHandler{
		componentBoggart: b,
		componentMQTT:    m,
	}
}

func (h *BindHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	q := r.URL().Query()

	var bindItem boggart.BindItem

	id := q.Get(":id")
	if id != "" {
		bindItem = h.componentBoggart.Bind(id)
		if bindItem == nil {
			h.NotFound(w, r)
			return
		}
	}

	switch q.Get(":action") {
	case "unregister":
		if bindItem.Type() == boggart.ComponentName {
			h.NotFound(w, r)
			return
		}

		h.actionDelete(w, r, bindItem)
		return

	case "readiness":
		h.actionProbe(w, r, bindItem, "readiness")
		return

	case "liveness":
		h.actionProbe(w, r, bindItem, "liveness")
		return

	case "logs":
		h.actionLogs(w, r, bindItem)
		return

	case "metrics":
		h.actionMetrics(w, r, bindItem)
		return

	case "mqtt":
		h.actionMQTT(w, r, bindItem)
		return

	case "tasks":
		h.actionTasks(w, r, bindItem)
		return

	case "":
		if bindItem != nil && bindItem.Type() == boggart.ComponentName {
			h.NotFound(w, r)
			return
		}

		h.actionCreateOrUpdate(w, r, bindItem)
		return

	default:
		h.NotFound(w, r)
		return
	}
}

func (h *BindHandler) registerByYAML(oldID string, code []byte) (bindItem boggart.BindItem, upgraded bool, err error) {
	bindParsed := &BindYAML{}

	err = yaml.Unmarshal(code, bindParsed)
	if err != nil {
		return nil, false, err
	}

	if bindParsed.Type == "" {
		return nil, false, errors.New("bind type is empty")
	}

	kind, err := boggart.GetBindType(bindParsed.Type)
	if err != nil {
		return nil, false, err
	}

	cfg, md, err := boggart.ValidateBindConfig(kind, bindParsed.Config)
	if err != nil {
		return nil, false, err
	}

	bind, err := kind.CreateBind(cfg)
	if err != nil {
		return nil, false, err
	}

	removeIDs := make([]string, 0, 2)

	// check new ID
	if bindParsed.ID != "" {
		removeIDs = append(removeIDs, bindParsed.ID)
	}

	// check old ID
	if oldID != "" && oldID != bindParsed.ID {
		removeIDs = append(removeIDs, oldID)
	}

	for _, id := range removeIDs {
		if bindExists := h.componentBoggart.Bind(id); bindExists != nil {
			upgraded = true

			if err := h.componentBoggart.UnregisterBindByID(id); err != nil {
				return nil, false, err
			}
		}
	}

	bindItem, err = h.componentBoggart.RegisterBind(bindParsed.ID, bind, bindParsed.Type, bindParsed.Description, bindParsed.Tags, cfg)

	if len(md.Unused) > 0 {
		if logger, ok := di.LoggerContainerBind(bind); ok {
			for _, field := range md.Unused {
				logger.Warn("Unused config field", "field", field)
			}
		}
	}

	return bindItem, upgraded, err
}

func (h *BindHandler) actionCreateOrUpdate(w *dashboard.Response, r *dashboard.Request, b boggart.BindItem) {
	var err error

	buf := bytes.NewBuffer(nil)
	vars := make(map[string]interface{})

	if r.IsPost() {
		code := r.Original().FormValue("yaml")
		if code == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		buf.WriteString(code)

		var (
			id       string
			bind     boggart.BindItem
			upgraded bool
		)

		if b != nil {
			id = b.ID()
		}

		if bind, upgraded, err = h.registerByYAML(id, buf.Bytes()); err == nil {
			if upgraded {
				r.Session().FlashBag().Info("Bind " + bind.ID() + " upgraded")
			} else {
				r.Session().FlashBag().Success("Bind register success with id " + bind.ID())
			}

			h.Redirect(r.URL().Path, http.StatusFound, w, r)

			return
		}
	} else {
		enc := yaml.NewEncoder(buf)
		defer enc.Close()

		if b == nil {
			bindYAML := &BindYAML{
				Description: "Description of new bind",
				Tags:        []string{"tag_label"},
				Config: map[string]interface{}{
					"config_key": "config_value",
				},
			}
			isAjax := false

			if typeName := r.URL().Query().Get("type"); typeName != "" && r.IsAjax() {
				bindYAML.Type = typeName
				isAjax = true
			} else {
				types := make([]string, 0)
				for typeName := range boggart.GetBindTypes() {
					types = append(types, typeName)
				}
				sort.Strings(types)
				vars["types"] = types

				if len(types) > 0 {
					bindYAML.Type = types[0]
				}
			}

			if bindYAML.Type != "" {
				if t, err := boggart.GetBindType(bindYAML.Type); err == nil {
					bindYAML.Config = t.Config()
				} else {
					bindYAML.Type = ""
				}
			}

			err = enc.Encode(bindYAML)

			if isAjax {
				if err == nil {
					_ = w.SendJSON(buf.String())
					return
				}

				h.InternalError(w, r, err)
				return
			}
		} else {
			err = enc.Encode(b)
		}
	}

	vars["yaml"] = buf.String()

	if err != nil {
		r.Session().FlashBag().Error(err.Error())
	}

	if b != nil {
		vars["bindId"] = b.ID()
	}

	h.Render(r.Context(), "bind", vars)
}

func (h *BindHandler) actionDelete(w *dashboard.Response, r *dashboard.Request, b boggart.BindItem) {
	if b == nil {
		h.NotFound(w, r)
		return
	}

	err := h.componentBoggart.UnregisterBindByID(b.ID())
	if err != nil {
		_ = w.SendJSON(boggart.NewResponseJSON().FailedError(err))

		return
	}

	_ = w.SendJSON(boggart.NewResponseJSON().Success(""))
}

func (h *BindHandler) actionProbe(w *dashboard.Response, r *dashboard.Request, b boggart.BindItem, t string) {
	if b == nil {
		h.NotFound(w, r)
		return
	}

	bindSupport, ok := di.ProbesContainerBind(b.Bind())

	if !ok {
		h.NotFound(w, r)
		return
	}

	var err error

	switch t {
	case "readiness":
		if bindSupport.ReadinessTaskID() == "" {
			h.NotFound(w, r)
			return
		}

		err = bindSupport.ReadinessCheck(r.Context())
	case "liveness":
		if bindSupport.LivenessTaskID() == "" {
			h.NotFound(w, r)
			return
		}

		err = bindSupport.LivenessCheck(r.Context())
	}

	if err != nil {
		_ = w.SendJSON(boggart.NewResponseJSON().FailedError(err))
	}

	_ = w.SendJSON(boggart.NewResponseJSON().Success(""))
}

func (h *BindHandler) actionLogs(w http.ResponseWriter, r *dashboard.Request, b boggart.BindItem) {
	bindSupport, ok := b.Bind().(di.LoggerContainerSupport)
	if !ok {
		h.NotFound(w, r)
		return
	}

	if r.IsPost() && r.URL().Query().Get("clean") == "1" {
		bindSupport.Clean()

		r.Session().FlashBag().Success("Logs cleaned")

		h.Redirect(r.URL().Path, http.StatusFound, w, r)

		return
	}

	type logView struct {
		Level   string
		Time    time.Time
		Message string
		Context string
	}

	records := bindSupport.LastRecords()
	response := make([]logView, len(records))

	for i, record := range bindSupport.LastRecords() {
		response[i].Level = record.Level.String()
		response[i].Time = record.Time
		response[i].Message = record.Message

		val := record.ContextMap()
		if len(val) == 0 {
			continue
		}

		buf := bytes.NewBuffer(nil)
		enc := yaml.NewEncoder(buf)

		err := enc.Encode(val)
		if err != nil {
			enc.Close()

			h.InternalError(w, r, err)

			return
		}

		enc.Close()

		response[i].Context = buf.String()
	}

	h.Render(r.Context(), "logs", map[string]interface{}{
		"bind": b,
		"logs": response,
	})
}

func (h *BindHandler) actionMetrics(w http.ResponseWriter, r *dashboard.Request, b boggart.BindItem) {
	bindSupport, ok := b.Bind().(di.MetricsContainerSupport)
	if !ok {
		h.NotFound(w, r)
		return
	}

	measures, err := bindSupport.Metrics().Gather()
	if err != nil {
		r.Session().FlashBag().Error(err.Error())
	}

	h.Render(r.Context(), "metrics", map[string]interface{}{
		"bind":     b,
		"measures": measures,
	})
}

func (h *BindHandler) actionMQTT(w http.ResponseWriter, r *dashboard.Request, b boggart.BindItem) {
	bindSupport, ok := b.Bind().(di.MQTTContainerSupport)
	if !ok {
		h.NotFound(w, r)
		return
	}

	type itemView struct {
		Topic    string
		Calls    uint64
		Datetime time.Time
		Payload  interface{}
	}

	publishes := bindSupport.MQTT().Publishes()
	publishesItems := make([]itemView, 0, len(publishes))

	for _, item := range h.componentMQTT.CacheItems() {
		for sent, count := range publishes {
			if item.Topic().String() == sent.String() {
				view := itemView{
					Topic:    sent.String(),
					Calls:    count,
					Datetime: item.Datetime(),
					Payload:  item.Payload(),
				}

				publishesItems = append(publishesItems, view)

				break
			}
		}
	}

	subscribers := bindSupport.MQTT().Subscribers()

	h.Render(r.Context(), "mqtt", map[string]interface{}{
		"bind":        b,
		"publishes":   publishesItems,
		"subscribers": subscribers,
	})
}

func (h *BindHandler) actionTasks(w *dashboard.Response, r *dashboard.Request, b boggart.BindItem) {
	bindSupport, ok := b.Bind().(di.WorkersContainerSupport)
	if !ok {
		h.NotFound(w, r)
		return
	}

	ctr := bindSupport.Workers()

	ids := ctr.TasksID()
	type taskView struct {
		Task      tasks.Task
		Meta      *tasks.Meta
		ShortName string
	}
	tasksView := make([]taskView, len(ids))
	var err error

	for i, id := range ids {
		tasksView[i].Task, tasksView[i].Meta, err = ctr.TaskByID(id[0])
		if err == nil {
			tasksView[i].ShortName = ctr.TaskShortName(tasksView[i].Task.Name())
		} else {
			tasksView[i].ShortName = ctr.TaskShortName(id[1])
		}
	}

	h.Render(r.Context(), "tasks", map[string]interface{}{
		"bind":  b,
		"tasks": tasksView,
	})
}
