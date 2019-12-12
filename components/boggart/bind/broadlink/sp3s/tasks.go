package sp3s

import (
	"context"

	"github.com/kihamo/go-workers"
	"go.uber.org/multierr"
)

func (b *Bind) Tasks() []workers.Task {
	taskUpdater := b.WrapTaskIsOnline(b.taskUpdater)
	taskUpdater.SetRepeats(-1)
	taskUpdater.SetRepeatInterval(b.config.UpdaterInterval)
	taskUpdater.SetName("updater-" + b.SerialNumber())

	return []workers.Task{
		taskUpdater,
	}
}

func (b *Bind) taskUpdater(ctx context.Context) error {
	state, err := b.State()
	if err != nil {
		return err
	}

	sn := b.SerialNumber()

	if e := b.MQTTPublishAsync(ctx, b.config.TopicState, state); e != nil {
		err = multierr.Append(err, e)
	}

	if power, e := b.Power(); e == nil {
		metricPower.With("serial_number", sn).Set(power)

		if e := b.MQTTPublishAsync(ctx, b.config.TopicPower, power); e != nil {
			err = multierr.Append(err, e)
		}
	} else {
		err = multierr.Append(err, e)
	}

	return err
}
