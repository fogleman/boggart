package internal

import (
	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/providers/mobile"
	"github.com/kihamo/snitch"
)

const (
	MetricMobileMegafonBalance                  = boggart.ComponentName + "_mobile_megafon_balance_rubles_total"
	MetricMobileMegafonUsedVoice                = boggart.ComponentName + "_mobile_megafon_used_voice_minutes"
	MetricMobileMegafonUsedSms                  = boggart.ComponentName + "_mobile_megafon_used_sms"
	MetricMobileMegafonUsedInternet             = boggart.ComponentName + "_mobile_megafon_used_internet_gigabytes"
	MetricMobileMegafonUsedInternetProlongation = boggart.ComponentName + "_mobile_megafon_used_internet_prolongation_gigabytes"
)

var (
	metricMobileMegafonBalance                  = snitch.NewGauge(MetricMobileMegafonBalance, "Megafon balance in rubles")
	metricMobileMegafonUsedVoice                = snitch.NewGauge(MetricMobileMegafonUsedVoice, "Megafon used voice in minutes")
	metricMobileMegafonUsedSms                  = snitch.NewGauge(MetricMobileMegafonUsedSms, "Megafon used sms")
	metricMobileMegafonUsedInternet             = snitch.NewGauge(MetricMobileMegafonUsedInternet, "Megafon used internet in GB")
	metricMobileMegafonUsedInternetProlongation = snitch.NewGauge(MetricMobileMegafonUsedInternetProlongation, "Megafon used internet prolongation in GB")
)

func (c *MetricsCollector) UpdaterMobile() error {
	megafonPhone := c.component.config.String(boggart.ConfigMobileMegafonPhone)
	megafonPassword := c.component.config.String(boggart.ConfigMobileMegafonPassword)

	if megafonPhone == "" || megafonPassword == "" {
		return nil
	}

	client := mobile.NewMegafon(megafonPhone, megafonPassword)

	value, err := client.Balance()
	if err != nil {
		// FIXME: logging
		return err
	}

	metricMobileMegafonBalance.With("phone", megafonPhone).Set(float64(value))

	remainders, err := client.Remainders()
	if err != nil {
		// FIXME: logging
		return err
	}

	metricMobileMegafonUsedVoice.With("phone", megafonPhone).Set(remainders.Voice)
	metricMobileMegafonUsedSms.With("phone", megafonPhone).Set(remainders.Sms)
	metricMobileMegafonUsedInternet.With("phone", megafonPhone).Set(remainders.Internet)
	metricMobileMegafonUsedInternetProlongation.With("phone", megafonPhone).Set(remainders.InternetProlongation)

	return nil
}

func (c *MetricsCollector) DescribeMobile(ch chan<- *snitch.Description) {
	metricMobileMegafonBalance.Describe(ch)
	metricMobileMegafonUsedVoice.Describe(ch)
	metricMobileMegafonUsedSms.Describe(ch)
	metricMobileMegafonUsedInternet.Describe(ch)
	metricMobileMegafonUsedInternetProlongation.Describe(ch)
}

func (c *MetricsCollector) CollectMobile(ch chan<- snitch.Metric) {
	metricMobileMegafonBalance.Collect(ch)
	metricMobileMegafonUsedVoice.Collect(ch)
	metricMobileMegafonUsedSms.Collect(ch)
	metricMobileMegafonUsedInternet.Collect(ch)
	metricMobileMegafonUsedInternetProlongation.Collect(ch)
}
