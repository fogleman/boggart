package google_home

import (
	"context"
)

func (b *Bind) ReadinessProbe(ctx context.Context) (err error) {
	_, err = b.GetSerialNumber(ctx)
	return err
}
