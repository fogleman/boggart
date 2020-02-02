package openhab

import (
	"net/http/httputil"

	"github.com/kihamo/boggart/components/boggart/di"
	"github.com/kihamo/boggart/providers/openhab"
)

type Bind struct {
	di.MetaBind
	di.LoggerBind
	di.ProbesBind

	config   *Config
	provider *openhab.Client
	proxy    *httputil.ReverseProxy
}