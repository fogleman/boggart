package handlers

import (
	"bytes"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/di"
	"github.com/kihamo/shadow/components/dashboard"
	"gopkg.in/yaml.v2"
)

// easyjson:json
type managerHandlerDevice struct {
	Id                string   `json:"id"`
	Type              string   `json:"type"`
	Description       string   `json:"description"`
	SerialNumber      string   `json:"serial_number"`
	MAC               string   `json:"mac"`
	Status            string   `json:"status"`
	Tasks             int      `json:"tasks"`
	MQTTPublishes     []string `json:"mqtt_publishes"`
	MQTTSubscribers   []string `json:"mqtt_subscribers"`
	Tags              []string `json:"tags"`
	Config            string   `json:"config"`
	HasWidget         bool     `json:"has_widget"`
	HasReadinessProbe bool     `json:"has_readiness_probe"`
	HasLivenessProbe  bool     `json:"has_liveness_probe"`
	LogsCount         int      `json:"logs_count"`
}

type ManagerHandler struct {
	dashboard.Handler

	component boggart.Component
}

func NewManagerHandler(component boggart.Component) *ManagerHandler {
	return &ManagerHandler{
		component: component,
	}
}

func (h *ManagerHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	if !r.IsAjax() {
		h.Render(r.Context(), "manager", map[string]interface{}{
			"device_types": boggart.GetBindTypes(),
		})
		return
	}

	entities := struct {
		Draw     int         `json:"draw"`
		Total    int         `json:"recordsTotal"`
		Filtered int         `json:"recordsFiltered"`
		Data     interface{} `json:"data"`
		Error    string      `json:"error,omitempty"`
	}{}
	entities.Draw = 1

	switch r.URL().Query().Get("entity") {
	case "devices":
		list := make([]managerHandlerDevice, 0)
		buf := bytes.NewBuffer(nil)
		enc := yaml.NewEncoder(buf)
		defer enc.Close()

		for _, bindItem := range h.component.BindItems() {
			buf.Reset()
			if err := enc.Encode(bindItem); err != nil {
				panic(err.Error())
			}

			item := managerHandlerDevice{
				Id:              bindItem.ID(),
				Type:            bindItem.Type(),
				Description:     bindItem.Description(),
				Status:          bindItem.Status().String(),
				Tags:            bindItem.Tags(),
				MQTTPublishes:   make([]string, 0),
				MQTTSubscribers: make([]string, 0),
				Config:          buf.String(),
			}

			if _, ok := bindItem.BindType().(boggart.BindTypeHasWidget); ok {
				item.HasWidget = ok
			}

			if bindSupport, ok := di.ProbesContainerBind(bindItem.Bind()); ok {
				item.HasReadinessProbe = bindSupport.Readiness() != nil
				item.HasLivenessProbe = bindSupport.Liveness() != nil
			}

			if bindSupport, ok := di.MetaContainerBind(bindItem.Bind()); ok {
				item.SerialNumber = bindSupport.SerialNumber()

				if mac := bindSupport.MAC(); mac != nil {
					item.MAC = mac.String()
				}
			}

			if bindSupport, ok := di.WorkersContainerBind(bindItem.Bind()); ok {
				item.Tasks = len(bindSupport.Tasks())
			}

			if bindSupport, ok := di.MQTTContainerBind(bindItem.Bind()); ok {
				for _, topic := range bindSupport.Publishes() {
					item.MQTTPublishes = append(item.MQTTPublishes, topic.String())
				}

				for _, topic := range bindSupport.Subscribers() {
					item.MQTTSubscribers = append(item.MQTTSubscribers, topic.Topic().String())
				}
			}

			if bindSupport, ok := bindItem.Bind().(di.LoggerContainerSupport); ok {
				item.LogsCount = len(bindSupport.LastRecords())
			}

			list = append(list, item)
		}

		entities.Data = list
		entities.Total = len(list)

	default:
		h.NotFound(w, r)
		return
	}

	entities.Filtered = entities.Total
	_ = w.SendJSON(entities)
}
