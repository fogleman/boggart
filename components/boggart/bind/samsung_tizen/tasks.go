package samsung_tizen

import (
	"context"
	"strings"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/go-workers"
	"github.com/kihamo/go-workers/task"
)

func (b *Bind) Tasks() []workers.Task {
	taskLiveness := task.NewFunctionTask(b.taskLiveness)
	taskLiveness.SetTimeout(b.livenessTimeout)
	taskLiveness.SetRepeats(-1)
	taskLiveness.SetRepeatInterval(b.livenessInterval)
	taskLiveness.SetName("bind-samsung-tizen-liveness")

	return []workers.Task{
		taskLiveness,
	}
}

func (b *Bind) taskLiveness(ctx context.Context) (interface{}, error) {
	info, err := b.client.Device(ctx)
	if err != nil {
		b.UpdateStatus(boggart.DeviceStatusOffline)
		return nil, nil
	}

	if b.Status() == boggart.DeviceStatusOnline {
		return nil, nil
	}

	if b.SerialNumber() == "" {
		parts := strings.Split(info.ID, ":")
		if len(parts) > 1 {
			b.SetSerialNumber(parts[1])
		} else {
			b.SetSerialNumber(info.ID)
		}

		b.mutex.Lock()
		b.mac = info.Device.WifiMac
		b.mutex.Unlock()

		sn := b.SerialNumber()
		b.MQTTPublishAsync(ctx, MQTTTopicDeviceID.Format(sn), 0, false, info.Device.ID)
		b.MQTTPublishAsync(ctx, MQTTTopicDeviceModelName.Format(sn), 0, false, info.Device.Name)
	}
	b.UpdateStatus(boggart.DeviceStatusOnline)

	return nil, nil
}