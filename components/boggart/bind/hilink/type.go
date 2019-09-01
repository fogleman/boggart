package hilink

import (
	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/atomic"
	"github.com/kihamo/boggart/components/boggart/protocols/swagger"
	"github.com/kihamo/boggart/components/boggart/providers/hilink"
)

type Type struct {
	boggart.BindTypeWidget
}

func (t Type) CreateBind(c interface{}) (_ boggart.Bind, err error) {
	config := c.(*Config)

	bind := &Bind{
		config:                    config,
		operator:                  atomic.NewString(),
		limitInternetTrafficIndex: atomic.NewInt64(),
		signalRSSI:                atomic.NewInt64Null(),
		signalRSRP:                atomic.NewInt64Null(),
		signalRSRQ:                atomic.NewInt64Null(),
		signalSINR:                atomic.NewInt64Null(),
		signalLevel:               atomic.NewInt64Null(),
	}

	l := swagger.NewLogger(
		func(message string) {
			bind.Logger().Info(message)
		},
		func(message string) {
			bind.Logger().Debug(message)
		})

	bind.client = hilink.New(config.Address.Host, config.Debug, l)

	return bind, nil
}
