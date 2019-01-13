package nut

import (
	"context"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/mqtt"
)

const (
	MQTTTopicVariable mqtt.Topic = boggart.ComponentName + "/ups/+/variable/+"
	MQTTTopicCommand  mqtt.Topic = boggart.ComponentName + "/ups/+/command"
)

/*
upscmd -l apcrs900g
Instant commands supported on UPS [apcrs900g]:

beeper.disable - Disable the UPS beeper
beeper.enable - Enable the UPS beeper
beeper.mute - Temporarily mute the UPS beeper
beeper.off - Obsolete (use beeper.disable or beeper.mute)
beeper.on - Obsolete (use beeper.enable)
load.off - Turn off the load immediately
load.off.delay - Turn off the load with a delay (seconds)
shutdown.reboot - Shut down the load briefly while rebooting the UPS
shutdown.stop - Stop a shutdown in progress
test.battery.start.deep - Start a deep battery test
test.battery.start.quick - Start a quick battery test
test.battery.stop - Stop the battery test
test.panel.start - Start testing the UPS panel
test.panel.stop - Stop a UPS panel test
*/

func (b *Bind) MQTTPublishes() []mqtt.Topic {
	return []mqtt.Topic{
		mqtt.Topic(MQTTTopicVariable.String()),
	}
}

func (b *Bind) MQTTSubscribers() []mqtt.Subscriber {
	return []mqtt.Subscriber{
		mqtt.NewSubscriber(MQTTTopicCommand.String(), 0, boggart.WrapMQTTSubscribeDeviceIsOnline(b, func(_ context.Context, _ mqtt.Component, message mqtt.Message) {
			if !boggart.CheckSerialNumberInMQTTTopic(b, message.Topic(), 2) {
				return
			}

			b.SendCommand(string(message.Payload()))
		})),
	}
}