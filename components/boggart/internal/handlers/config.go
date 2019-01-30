package handlers

import (
	"bytes"
	"io"
	"strconv"
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/internal/manager"
	"github.com/kihamo/shadow/components/dashboard"
	"gopkg.in/yaml.v2"
)

type ConfigHandler struct {
	dashboard.Handler

	manager   *manager.Manager
	component boggart.Component
}

func NewConfigHandler(manager *manager.Manager, component boggart.Component) *ConfigHandler {
	return &ConfigHandler{
		manager:   manager,
		component: component,
	}
}

func (h *ConfigHandler) ServeHTTP(w *dashboard.Response, r *dashboard.Request) {
	q := r.URL().Query()
	action := q.Get(":action")

	if action == "reload" && r.IsPost() {
		type response struct {
			Result  string `json:"result"`
			Message string `json:"message,omitempty"`
		}

		// reload by ID
		if id := q.Get("id"); id != "" {
			if err := h.component.ReloadConfigByID(id); err != nil {
				w.SendJSON(response{
					Result:  "failed",
					Message: err.Error(),
				})
			} else {
				w.SendJSON(response{
					Result:  "success",
					Message: "Bind " + id + " reloaded from file",
				})
			}

			return
		}

		// reload all
		if loaded, err := h.component.ReloadConfig(); err != nil {
			w.SendJSON(response{
				Result:  "failed",
				Message: err.Error(),
			})
		} else {
			w.SendJSON(response{
				Result:  "success",
				Message: "Loaded " + strconv.FormatInt(int64(loaded), 10) + " binds",
			})
		}

		return
	}

	if action != "download" && action != "modal" && action != "view" {
		h.NotFound(w, r)
	}

	buf := bytes.NewBuffer(nil)
	enc := yaml.NewEncoder(buf)
	defer enc.Close()

	if err := enc.Encode(h.manager.BindItems()); err != nil {
		panic(err.Error())
	}

	switch action {
	case "download":
		w.Header().Set("Content-Length", strconv.FormatInt(int64(buf.Len()), 10))
		w.Header().Set("Content-Disposition", "attachment; filename=\""+time.Now().Format("20060102150405_config.yaml")+"\"")
		w.Header().Set("Content-Type", "text/x-yaml")

		if _, err := io.Copy(w, buf); err != nil {
			panic(err.Error())
		}

	case "modal":
		h.RenderLayout(r.Context(), "config", "simple", map[string]interface{}{
			"yaml":  buf.String(),
			"modal": true,
		})

	case "view":
		h.Render(r.Context(), "config", map[string]interface{}{
			"yaml":  buf.String(),
			"modal": false,
		})
	}
}
