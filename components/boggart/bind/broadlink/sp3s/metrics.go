package sp3s

import (
	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/snitch"
)

var (
	metricPower = snitch.NewGauge(boggart.ComponentName+"_bind_broadlink_sp3s_power_watt", "Broadlink SP3S socket current power")
)

func (b *Bind) Describe(ch chan<- *snitch.Description) {
	serialNumber := b.SerialNumber()

	metricPower.With("serial_number", serialNumber).Describe(ch)
}

func (b *Bind) Collect(ch chan<- snitch.Metric) {
	serialNumber := b.SerialNumber()

	metricPower.With("serial_number", serialNumber).Collect(ch)
}