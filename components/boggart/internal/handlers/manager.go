package handlers

import (
	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/config_generators"
	"github.com/kihamo/boggart/components/boggart/di"
	"github.com/kihamo/boggart/components/boggart/tasks"
	"github.com/kihamo/shadow/components/dashboard"
	"go.uber.org/zap/zapcore"
)

// easyjson:json
type managerHandlerDeviceTask struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Registered     bool   `json:"registered"`
	CustomSchedule bool   `json:"custom_schedule"`
}

// easyjson:json
type managerHandlerDevice struct {
	Tasks                    []managerHandlerDeviceTask `json:"tasks"`
	Tags                     []string                   `json:"tags"`
	ConfigGenerators         []string                   `json:"config_generators"`
	ID                       string                     `json:"id"`
	Type                     string                     `json:"type"`
	Description              string                     `json:"description"`
	SerialNumber             string                     `json:"serial_number"`
	MAC                      string                     `json:"mac"`
	Status                   string                     `json:"status"`
	MetricsDescriptionsCount uint64                     `json:"metrics_descriptions_count"`
	MetricsCollectCount      uint64                     `json:"metrics_collect_count"`
	MetricsEmptyCount        uint64                     `json:"metrics_empty_count"`
	MQTTPublishes            int                        `json:"mqtt_publishes"`
	MQTTSubscribers          int                        `json:"mqtt_subscribers"`
	LogsCount                int                        `json:"logs_count"`
	HasMetrics               bool                       `json:"has_metrics"`
	HasWidget                bool                       `json:"has_widget"`
	LogsMaxLevel             zapcore.Level              `json:"logs_max_level"`
}

type ManagerHandler struct {
	dashboard.Handler

	component boggart.Component
	manager   *tasks.Manager
}

func NewManagerHandler(component boggart.Component, manager *tasks.Manager) *ManagerHandler {
	return &ManagerHandler{
		component: component,
		manager:   manager,
	}
}

func (h *ManagerHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	if !r.IsAjax() {
		h.Render(r.Context(), "manager", map[string]interface{}{
			"device_types": boggart.GetBindTypes(),
		})

		return
	}

	entities := boggart.NewResponseDataTable()

	switch r.URL().Query().Get("entity") {
	case "devices":
		list := make([]managerHandlerDevice, 0)

		for _, bindItem := range h.component.BindItems() {
			bind := bindItem.Bind()
			item := managerHandlerDevice{
				ID:               bindItem.ID(),
				Type:             bindItem.Type(),
				Description:      bindItem.Description(),
				Status:           bindItem.Status().String(),
				Tags:             bindItem.Tags(),
				LogsMaxLevel:     zapcore.DebugLevel,
				ConfigGenerators: make([]string, 0),
			}

			if bindSupport, ok := di.WidgetContainerBind(bind); ok {
				item.HasWidget = ok && bindSupport.HandleAllowed()
			}

			if bindSupport, ok := di.MetaContainerBind(bind); ok {
				item.SerialNumber = bindSupport.SerialNumber()

				if mac := bindSupport.MAC(); mac != nil {
					item.MAC = mac.String()
				}
			}

			if bindSupport, ok := di.WorkersContainerBind(bind); ok {
				ids := bindSupport.TasksID()
				item.Tasks = make([]managerHandlerDeviceTask, len(ids))
				for i, id := range ids {
					item.Tasks[i].ID = id[0]
					item.Tasks[i].Name = bindSupport.TaskShortName(id[1])

					t, _ := h.manager.Task(id[0])
					item.Tasks[i].Registered = t != nil
				}
			} else if bindSupport, ok := di.ProbesContainerBind(bind); ok {
				item.Tasks = make([]managerHandlerDeviceTask, 0, 2)

				if taskID := bindSupport.ReadinessTaskID(); taskID != "" {
					itemView := managerHandlerDeviceTask{
						ID:   taskID,
						Name: di.ProbesConfigReadinessDefaultName,
					}

					t, _ := h.manager.Task(taskID)
					itemView.Registered = t != nil

					_, ok = bind.(di.BindHasReadinessProbeSchedule)
					itemView.CustomSchedule = ok

					item.Tasks = append(item.Tasks, itemView)
				}

				if taskID := bindSupport.LivenessTaskID(); taskID != "" {
					itemView := managerHandlerDeviceTask{
						ID:   taskID,
						Name: di.ProbesConfigLivenessDefaultName,
					}

					t, _ := h.manager.Task(taskID)
					itemView.Registered = t != nil

					_, ok = bind.(di.BindHasLivenessProbeSchedule)
					itemView.CustomSchedule = ok

					item.Tasks = append(item.Tasks, itemView)
				}
			}

			if bindSupport, ok := di.MetricsContainerBind(bind); ok {
				item.HasMetrics = true
				item.MetricsDescriptionsCount = bindSupport.DescriptionsCount()
				item.MetricsCollectCount = bindSupport.CollectCount()
				item.MetricsEmptyCount = bindSupport.EmptyCount()
			}

			if bindSupport, ok := di.MQTTContainerBind(bind); ok {
				item.MQTTSubscribers = len(bindSupport.Subscribers())
				item.MQTTPublishes = len(bindSupport.Publishes())
			}

			if bindSupport, ok := bind.(di.LoggerContainerSupport); ok {
				records := bindSupport.LastRecords()
				item.LogsCount = len(records)

				for _, r := range records {
					if r.Level > item.LogsMaxLevel {
						item.LogsMaxLevel = r.Level
					}
				}
			}

			if _, ok := bind.(generators.HasGeneratorOpenHab); ok {
				item.ConfigGenerators = append(item.ConfigGenerators, "openhab")
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
