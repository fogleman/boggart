package internal

import (
	"github.com/kihamo/boggart/components/storage"
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/boggart/components/voice/internal/handlers"
	"github.com/kihamo/shadow/components/dashboard"
)

func (c *Component) DashboardTemplates() *assetfs.AssetFS {
	return dashboard.TemplatesFromAssetFS(c)
}

func (c *Component) DashboardMenu() dashboard.Menu {
	routes := c.DashboardRoutes()

	return dashboard.NewMenu("Voice").WithRoute(routes[0]).WithIcon("microphone")
}

func (c *Component) DashboardRoutes() []dashboard.Route {
	if c.routes == nil {
		<-c.application.ReadyComponent(storage.ComponentName)

		c.routes = []dashboard.Route{
			dashboard.NewRoute("/"+c.Name()+"/message/", &handlers.MessageHandler{
				Component: c,
			}).
				WithMethods([]string{http.MethodGet, http.MethodPost}).
				WithAuth(true),
			dashboard.NewRoute("/"+c.Name()+"/file/:message", &handlers.FileHandler{
				Voice:   c,
				Storage: c.application.GetComponent(storage.ComponentName).(storage.Component),
			}).
				WithMethods([]string{http.MethodGet}),
		}
	}

	return c.routes
}
