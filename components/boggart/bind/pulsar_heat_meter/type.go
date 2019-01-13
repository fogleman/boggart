package pulsar_heat_meter

import (
	"encoding/hex"
	"math"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/protocols/rs485"
	"github.com/kihamo/boggart/components/boggart/providers/pulsar"
)

type Type struct{}

func (t Type) CreateBind(c interface{}) (boggart.DeviceBind, error) {
	config := c.(*Config)

	var err error
	conn := rs485.GetConnection(config.RS485Address, config.RS485Timeout)

	var deviceAddress []byte
	if config.Address == "" {
		deviceAddress, err = pulsar.DeviceAddress(conn)
	} else {
		deviceAddress, err = hex.DecodeString(config.Address)
	}

	if err != nil {
		return nil, err
	}

	device := &Bind{
		config:   config,
		provider: pulsar.NewHeatMeter(deviceAddress, conn),

		temperatureIn:    math.MaxUint64,
		temperatureOut:   math.MaxUint64,
		temperatureDelta: math.MaxUint64,
		energy:           math.MaxUint64,
		consumption:      math.MaxUint64,
		capacity:         math.MaxUint64,
		power:            math.MaxUint64,
		input1:           math.MaxUint64,
		input2:           math.MaxUint64,
		input3:           math.MaxUint64,
		input4:           math.MaxUint64,

		updaterInterval: config.UpdaterInterval,
	}
	device.Init()
	device.SetSerialNumber(hex.EncodeToString(deviceAddress))

	return device, nil
}
