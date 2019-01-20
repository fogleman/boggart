package alsa

import (
	"bytes"
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/mqtt"
)

const (
	MQTTPrefix = boggart.ComponentName + "/alsa/+/"

	MQTTSubscribeTopicVolume    mqtt.Topic = MQTTPrefix + "volume"
	MQTTSubscribeTopicMute      mqtt.Topic = MQTTPrefix + "mute"
	MQTTSubscribeTopicPause     mqtt.Topic = MQTTPrefix + "pause"
	MQTTSubscribeTopicStop      mqtt.Topic = MQTTPrefix + "stop"
	MQTTSubscribeTopicPlay      mqtt.Topic = MQTTPrefix + "play"
	MQTTSubscribeTopicResume    mqtt.Topic = MQTTPrefix + "resume"
	MQTTSubscribeTopicAction    mqtt.Topic = MQTTPrefix + "action"
	MQTTPublishTopicStateStatus mqtt.Topic = MQTTPrefix + "state/status"
	MQTTPublishTopicStateVolume mqtt.Topic = MQTTPrefix + "state/volume"
	MQTTPublishTopicStateMute   mqtt.Topic = MQTTPrefix + "state/mute"
)

func (b *Bind) MQTTPublishes() []mqtt.Topic {
	sn := mqtt.NameReplace(b.SerialNumber())

	return []mqtt.Topic{
		mqtt.Topic(MQTTPublishTopicStateStatus.Format(sn)),
		mqtt.Topic(MQTTPublishTopicStateVolume.Format(sn)),
		mqtt.Topic(MQTTPublishTopicStateMute.Format(sn)),
	}
}

func (b *Bind) MQTTSubscribers() []mqtt.Subscriber {
	sn := mqtt.NameReplace(b.SerialNumber())

	return []mqtt.Subscriber{
		mqtt.NewSubscriber(MQTTSubscribeTopicVolume.Format(sn), 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b, func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
			volume, err := strconv.ParseInt(string(message.Payload()), 10, 64)
			if err != nil {
				return err
			}

			return b.SetVolume(volume)
		})),
		mqtt.NewSubscriber(MQTTSubscribeTopicMute.Format(sn), 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b, func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
			return b.SetMute(bytes.Equal(message.Payload(), []byte(`1`)))
		})),
		mqtt.NewSubscriber(MQTTSubscribeTopicPause.Format(sn), 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b, func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
			return b.Pause()
		})),
		mqtt.NewSubscriber(MQTTSubscribeTopicStop.Format(sn), 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b, func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
			return b.Stop()
		})),
		mqtt.NewSubscriber(MQTTSubscribeTopicPlay.Format(sn), 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b, func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
			if u := string(message.Payload()); u != "" {
				return b.PlayFromURL(u)
			}

			return b.Play()
		})),
		mqtt.NewSubscriber(MQTTSubscribeTopicResume.Format(sn), 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b, func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
			if u := string(message.Payload()); u != "" {
				return b.PlayFromURL(u)
			}

			return b.Play()
		})),
		mqtt.NewSubscriber(MQTTSubscribeTopicAction.Format(sn), 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b, func(ctx context.Context, _ mqtt.Component, message mqtt.Message) error {
			action := string(message.Payload())

			switch strings.ToLower(action) {
			case "stop":
				return b.Stop()

			case "pause":
				return b.Pause()

			case "play", "resume":
				return b.Play()
			}

			return errors.New("unknown action " + action)
		})),
	}
}
