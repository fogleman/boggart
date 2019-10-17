package octoprint

import (
	"github.com/kihamo/boggart/components/mqtt"
)

func (b *Bind) MQTTPublishes() []mqtt.Topic {
	sn := b.config.Address.Host

	return []mqtt.Topic{
		b.config.TopicState.Format(sn),
		b.config.TopicStateBedTemperatureActual.Format(sn),
		b.config.TopicStateBedTemperatureOffset.Format(sn),
		b.config.TopicStateBedTemperatureTarget.Format(sn),
		b.config.TopicStateTool0TemperatureActual.Format(sn),
		b.config.TopicStateTool0TemperatureOffset.Format(sn),
		b.config.TopicStateTool0TemperatureTarget.Format(sn),
	}
}
