package scale

import (
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/mqtt"
	"github.com/kihamo/boggart/providers/xiaomi/scale"
)

type Config struct {
	MAC             boggart.HardwareAddr `valid:",required"`
	Device          string
	UpdaterInterval time.Duration `mapstructure:"updater_interval" yaml:"updater_interval"`
	CaptureDuration time.Duration `mapstructure:"capture_interval" yaml:"capture_interval"`
	TopicWeight     mqtt.Topic    `mapstructure:"topic_weight" yaml:"topic_weight"`
}

func (Type) Config() interface{} {
	var prefix mqtt.Topic = boggart.ComponentName + "/xiaomi/scale/+/"

	return &Config{
		Device:          scale.DefaultDevice,
		UpdaterInterval: time.Minute,
		CaptureDuration: time.Second * 10,
		TopicWeight:     prefix + "weight",
	}
}
