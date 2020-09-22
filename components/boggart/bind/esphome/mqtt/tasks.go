package mqtt

import (
	"context"
	"net/http"
	"net/url"

	"github.com/kihamo/go-workers"
	"github.com/kihamo/snitch"
	"github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

func (b *Bind) Tasks() []workers.Task {
	if b.config.TopicIPAddressSensor == "" {
		return nil
	}

	taskImportMetrics := b.Workers().WrapTaskIsOnline(b.taskImportMetrics)
	taskImportMetrics.SetRepeats(-1)
	taskImportMetrics.SetRepeatInterval(b.config.ImportMetricsInterval)
	taskImportMetrics.SetName("import-metrics")

	return []workers.Task{
		taskImportMetrics,
	}
}

func (b *Bind) taskImportMetrics(ctx context.Context) error {
	mac := b.Meta().MACAsString()
	if mac == "" {
		return nil
	}

	b.ipMutex.RLock()
	ip := b.ip
	b.ipMutex.RUnlock()

	if ip == nil {
		return nil
	}

	u := &url.URL{
		Scheme: "http",
		Host:   ip.String(),
		Path:   "/metrics",
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	parser := expfmt.TextParser{}
	m, err := parser.TextToMetricFamilies(response.Body)
	if err != nil {
		return err
	}

	for key, outerMetric := range m {
		t := outerMetric.GetType()
		if t != io_prometheus_client.MetricType_GAUGE {
			b.Logger().Warn("Unknown metric type", "key", key, "type", t.String())
			continue
		}

		var innerMetric snitch.Gauge

		switch key {
		case "esphome_sensor_value":
			innerMetric = metricSensorValue
		case "esphome_sensor_failed":
			innerMetric = metricSensorFailed
		case "esphome_binary_sensor_value":
			innerMetric = metricBinarySensorValue
		case "esphome_binary_sensor_failed":
			innerMetric = metricBinarySensorFailed
		case "esphome_fan_value":
			innerMetric = metricFanValue
		case "esphome_fan_failed":
			innerMetric = metricFanFailed
		case "esphome_fan_speed":
			innerMetric = metricFanSpeed
		case "esphome_fan_oscillation":
			innerMetric = metricFanOscillation
		case "esphome_light_state":
			innerMetric = metricLightState
		case "esphome_light_color":
			innerMetric = metricLightColor
		case "esphome_light_effect_active":
			innerMetric = metricLightEffectActive
		case "esphome_cover_value":
			innerMetric = metricCoverValue
		case "esphome_cover_failed":
			innerMetric = metricCoverFailed
		case "esphome_switch_value":
			innerMetric = metricSwitchValue
		case "esphome_switch_failed":
			innerMetric = metricSwitchFailed
		}

		if innerMetric == nil {
			b.Logger().Warn("Unknown metric key", "key", key, "type", t.String())
			continue
		}

		innerMetric = innerMetric.With("mac", mac)

		for _, subMetric := range outerMetric.GetMetric() {
			gauge := subMetric.GetGauge()
			if gauge == nil {
				b.Logger().Warn("Gauge metric is empty", "key", key)
				continue
			}

			outerLabels := subMetric.GetLabel()
			innerLabels := make([]string, 0, len(outerLabels)*2)

			for _, label := range outerLabels {
				innerLabels = append(innerLabels, label.GetName(), label.GetValue())
			}

			innerMetric.With(innerLabels...).Set(gauge.GetValue())
		}
	}

	return nil
}
