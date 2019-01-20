package samsung_tizen

import (
	"bytes"
	"context"
	"errors"

	"github.com/ghthor/gowol"
	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/providers/samsung/tv"
	"github.com/kihamo/boggart/components/mqtt"
)

const (
	MQTTSubscribeTopicPower         mqtt.Topic = boggart.ComponentName + "/tv/+/power"
	MQTTSubscribeTopicKey           mqtt.Topic = boggart.ComponentName + "/tv/+/key"
	MQTTPublishTopicDeviceID        mqtt.Topic = boggart.ComponentName + "/tv/+/device/id"
	MQTTPublishTopicDeviceModelName mqtt.Topic = boggart.ComponentName + "/tv/+/device/model-name"
)

func (b *Bind) MQTTPublishes() []mqtt.Topic {
	return []mqtt.Topic{
		MQTTPublishTopicDeviceID,
		MQTTPublishTopicDeviceModelName,
	}
}

func (b *Bind) MQTTSubscribers() []mqtt.Subscriber {
	return []mqtt.Subscriber{
		mqtt.NewSubscriber(MQTTSubscribeTopicPower.String(), 0, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 2) {
				return nil
			}

			if bytes.Equal(message.Payload(), []byte(`1`)) {
				b.mutex.RLock()
				mac := b.mac
				b.mutex.RUnlock()

				return wol.MagicWake(mac, "255.255.255.255")
			}

			if b.Status() != boggart.BindStatusOnline {
				return errors.New("bind isn't online")
			}

			return b.client.SendCommand(tv.KeyPower)
		}),
		mqtt.NewSubscriber(MQTTSubscribeTopicKey.String(), 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b.Status, func(_ context.Context, _ mqtt.Component, message mqtt.Message) error {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 2) {
				return nil
			}

			return b.client.SendCommand(string(message.Payload()))
		})),
	}
}
