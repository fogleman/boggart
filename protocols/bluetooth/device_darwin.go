// +build darwin

package bluetooth

import (
	"github.com/go-ble/ble"
)

// DefaultDevice ...
func NewDevice(opts ...ble.Option) (d ble.Device, err error) {
	//return darwin.NewDevice(opts...)
	return NewDummyDevice(opts...)
}
