package internal

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/boggart/components/boggart/internal/handlers"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/messengers"
)

func (c *Component) DashboardTemplates() *assetfs.AssetFS {
	return dashboard.TemplatesFromAssetFS(c)
}

func (c *Component) DashboardMenu() dashboard.Menu {
	routes := c.DashboardRoutes()

	return dashboard.NewMenu("Smart home").
		WithIcon("home").
		WithChild(dashboard.NewMenu("Manager").WithRoute(routes[1])).
		WithChild(dashboard.NewMenu("Config YAML").WithUrl("/" + c.Name() + "/config/view"))
}

func (c *Component) DashboardRoutes() []dashboard.Route {
	if c.routes == nil {
		<-c.application.ReadyComponent(c.Name())

		cameraHandler := &handlers.CameraHandler{
			DevicesManager: c.devicesManager,
		}

		m := c.application.GetComponent(messengers.ComponentName)
		if m != nil {
			<-c.application.ReadyComponent(messengers.ComponentName)
			cameraHandler.Messengers = m.(messengers.Component)
		}

		c.routes = []dashboard.Route{
			dashboard.RouteFromAssetFS(c),
			dashboard.NewRoute("/"+c.Name()+"/manager/", handlers.NewManagerIndexHandler(c.devicesManager, c.listenersManager)).
				WithMethods([]string{http.MethodGet}).
				WithAuth(true),
			dashboard.NewRoute("/"+c.Name()+"/manager/:id/unregister", handlers.NewManagerUnregisterHandler(c.devicesManager)).
				WithMethods([]string{http.MethodPost}).
				WithAuth(true),
			dashboard.NewRoute("/"+c.Name()+"/camera/:id/:channel", cameraHandler).
				WithMethods([]string{http.MethodGet, http.MethodPost}),
			dashboard.NewRoute("/"+c.Name()+"/config/:action", handlers.NewConfigHandler(c.devicesManager)).
				WithMethods([]string{http.MethodGet}),
		}
	}

	return c.routes
}
