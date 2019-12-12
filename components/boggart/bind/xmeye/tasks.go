package xmeye

import (
	"context"
	"errors"
	"strconv"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/go-workers"
	"github.com/kihamo/go-workers/task"
	"go.uber.org/multierr"
)

func (b *Bind) Tasks() []workers.Task {
	taskLiveness := task.NewFunctionTask(b.taskLiveness)
	taskLiveness.SetTimeout(b.config.LivenessTimeout)
	taskLiveness.SetRepeats(-1)
	taskLiveness.SetRepeatInterval(b.config.LivenessInterval)
	taskLiveness.SetName("liveness-" + b.config.Address.Host)

	taskState := b.WrapTaskIsOnline(b.taskUpdater)
	taskState.SetTimeout(b.config.UpdaterTimeout)
	taskState.SetRepeats(-1)
	taskState.SetRepeatInterval(b.config.UpdaterInterval)
	taskState.SetName("updater-" + b.config.Address.Host)

	tasks := []workers.Task{
		taskLiveness,
		taskState,
	}

	return tasks
}

func (b *Bind) taskLiveness(ctx context.Context) (interface{}, error) {
	client, err := b.client()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	info, err := client.SystemInfo(ctx)
	if err != nil {
		b.UpdateStatus(boggart.BindStatusOffline)
		return nil, err
	}

	if info.SerialNo == "" {
		b.UpdateStatus(boggart.BindStatusOffline)
		return nil, errors.New("device returns empty serial number")
	}

	if b.SerialNumber() == "" {
		b.SetSerialNumber(info.SerialNo)

		if b.config.AlarmStreamingEnabled {
			if err = b.startAlarmStreaming(); err != nil {
				return nil, err
			}
		}

		if e := b.MQTTPublishAsync(ctx, b.config.TopicStateModel.Format(info.SerialNo), info.HardWare); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTTPublishAsync(ctx, b.config.TopicStateFirmwareVersion.Format(info.SerialNo), info.SoftWareVersion); e != nil {
			err = multierr.Append(err, e)
		}

		if e := b.MQTTPublishAsync(ctx, b.config.TopicStateFirmwareReleasedDate.Format(info.SerialNo), info.BuildTime); e != nil {
			err = multierr.Append(err, e)
		}
	}

	b.UpdateStatus(boggart.BindStatusOnline)

	return nil, err
}

func (b *Bind) taskUpdater(ctx context.Context) (err error) {
	if !b.IsStatusOnline() {
		return nil
	}

	client, err := b.client()
	if err != nil {
		return err
	}
	defer client.Close()

	sn := b.SerialNumber()

	storage, _ := client.StorageInfo(ctx)
	for _, s := range storage {
		for _, p := range s.Partition {
			if p.IsCurrent {
				name := strconv.FormatUint(p.LogicSerialNo, 10)

				// TODO:
				if e := b.MQTTPublishAsync(ctx, b.config.TopicStateHDDCapacity.Format(sn, p.LogicSerialNo), uint64(p.TotalSpace)*MB); e != nil {
					err = multierr.Append(err, e)
				}

				// TODO:
				if e := b.MQTTPublishAsync(ctx, b.config.TopicStateHDDUsage.Format(sn, p.LogicSerialNo), uint64(p.TotalSpace-p.RemainSpace)*MB); e != nil {
					err = multierr.Append(err, e)
				}
				metricStorageUsage.With("serial_number", sn).With("name", name).Set(float64(uint64(p.TotalSpace-p.RemainSpace) * MB))

				// TODO:
				if e := b.MQTTPublishAsync(ctx, b.config.TopicStateHDDFree.Format(sn, p.LogicSerialNo), uint64(p.RemainSpace)*MB); e != nil {
					err = multierr.Append(err, e)
				}
				metricStorageAvailable.With("serial_number", sn).With("name", name).Set(float64(uint64(p.RemainSpace) * MB))
			}
		}
	}

	return err
}
