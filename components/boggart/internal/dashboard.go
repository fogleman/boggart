package internal

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/internal/handlers"
	"github.com/kihamo/shadow/components/dashboard"
)

func (c *Component) GetTemplates() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "templates",
	}
}

func (c *Component) GetDashboardMenu() dashboard.Menu {
	routes := c.GetDashboardRoutes()
	menus := []dashboard.Menu{
		dashboard.NewMenuWithRoute("Dashboard", routes[0], "", nil, nil),
		dashboard.NewMenuWithRoute("Detect", routes[1], "", nil, nil),
		dashboard.NewMenuWithRoute("Devices", routes[2], "", nil, nil),
	}

	if u := c.config.GetString(boggart.ConfigMonitoringExternalURL); u != "" {
		menus = append(menus, dashboard.NewMenuWithUrl("Monitoring", u, "", nil, nil))
	}

	return dashboard.NewMenuWithUrl("Boggart", "/"+c.GetName()+"/", "home", menus, nil)
}

func (c *Component) GetDashboardRoutes() []dashboard.Route {
	if c.routes == nil {
		c.routes = []dashboard.Route{
			dashboard.NewRoute(
				c.GetName(),
				[]string{http.MethodGet},
				"/"+c.GetName()+"/",
				&handlers.IndexHandler{
					Config:    c.config,
					Component: c,
				},
				"",
				true),
			dashboard.NewRoute(
				c.GetName(),
				[]string{http.MethodGet},
				"/"+c.GetName()+"/detect/",
				&handlers.DetectHandler{},
				"",
				true),
			dashboard.NewRoute(
				c.GetName(),
				[]string{http.MethodGet},
				"/"+c.GetName()+"/devices/",
				&handlers.DevicesHandler{},
				"",
				true),
			dashboard.NewRoute(
				c.GetName(),
				[]string{http.MethodGet},
				"/"+c.GetName()+"/hikvision/:place/:action/",
				&handlers.HikvisionHandler{
					Config: c.config,
				},
				"",
				true),
		}
	}

	return c.routes
}
