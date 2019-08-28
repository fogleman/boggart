package hilink

import (
	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/snitch"
)

var (
	metricBalance              = snitch.NewGauge(boggart.ComponentName+"_bind_hilink_balance_rubles", "HiLink balance in rubles")
	metricLimitInternetTraffic = snitch.NewGauge(boggart.ComponentName+"_bind_hilink_limit_internet_traffic_bytes", "HiLink limit internet traffic in bytes")
)

func (b *Bind) Describe(ch chan<- *snitch.Description) {
	sn := b.SerialNumber()
	if sn == "" {
		return
	}

	metricBalance.With("serial_number", sn).Describe(ch)
	metricLimitInternetTraffic.With("serial_number", sn).Describe(ch)
}

func (b *Bind) Collect(ch chan<- snitch.Metric) {
	sn := b.SerialNumber()
	if sn == "" {
		return
	}

	metricBalance.With("serial_number", sn).Collect(ch)
	metricLimitInternetTraffic.With("serial_number", sn).Collect(ch)
}