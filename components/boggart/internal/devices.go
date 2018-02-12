package internal

import (
	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/devices"
	"github.com/kihamo/boggart/components/boggart/providers/hikvision"
	"github.com/kihamo/boggart/components/boggart/providers/mobile"
)

func (c *Component) initVideoRecorders() {
	isapi := hikvision.NewISAPI(
		c.config.String(boggart.ConfigVideoRecorderHikVisionHost),
		c.config.Int64(boggart.ConfigVideoRecorderHikVisionPort),
		c.config.String(boggart.ConfigVideoRecorderHikVisionUsername),
		c.config.String(boggart.ConfigVideoRecorderHikVisionPassword))

	device, err := devices.NewVideoRecorderHikVision(isapi)
	if err != nil {
		c.logger.Error("Init video recorder failed", map[string]interface{}{
			"error":    err.Error(),
			"host":     c.config.String(boggart.ConfigVideoRecorderHikVisionHost),
			"port":     c.config.String(boggart.ConfigVideoRecorderHikVisionPort),
			"username": c.config.String(boggart.ConfigVideoRecorderHikVisionUsername),
		})
		return
	}

	if c.config.Bool(boggart.ConfigVideoRecorderHikVisionEnabled) {
		device.Enable()
	} else {
		device.Disable()
	}

	c.devices.Register(boggart.DeviceVideoRecorderID, device)
}

func (c *Component) initCameras() {
	isapi := hikvision.NewISAPI(
		c.config.String(boggart.ConfigCameraHikVisionHallHost),
		c.config.Int64(boggart.ConfigCameraHikVisionHallPort),
		c.config.String(boggart.ConfigCameraHikVisionHallUsername),
		c.config.String(boggart.ConfigCameraHikVisionHallPassword))

	device, err := devices.NewCameraHikVision(isapi, c.config.Uint64(boggart.ConfigCameraHikVisionHallStreamingChannel))
	if err != nil {
		c.logger.Error("Init camera failed", map[string]interface{}{
			"error":    err.Error(),
			"host":     c.config.String(boggart.ConfigCameraHikVisionHallHost),
			"port":     c.config.String(boggart.ConfigCameraHikVisionHallPort),
			"username": c.config.String(boggart.ConfigCameraHikVisionHallUsername),
		})
		return
	} else {
		if c.config.Bool(boggart.ConfigCameraHikVisionHallEnabled) {
			device.Enable()
		} else {
			device.Disable()
		}

		c.devices.Register(boggart.DeviceCameraHallID, device)
	}

	isapi = hikvision.NewISAPI(
		c.config.String(boggart.ConfigCameraHikVisionStreetHost),
		c.config.Int64(boggart.ConfigCameraHikVisionStreetPort),
		c.config.String(boggart.ConfigCameraHikVisionStreetUsername),
		c.config.String(boggart.ConfigCameraHikVisionStreetPassword))

	device, err = devices.NewCameraHikVision(isapi, c.config.Uint64(boggart.ConfigCameraHikVisionStreetStreamingChannel))
	if err != nil {
		c.logger.Error("Init camera failed", map[string]interface{}{
			"error":    err.Error(),
			"host":     c.config.String(boggart.ConfigCameraHikVisionStreetHost),
			"port":     c.config.String(boggart.ConfigCameraHikVisionStreetPort),
			"username": c.config.String(boggart.ConfigCameraHikVisionStreetUsername),
		})
		return
	} else {
		if c.config.Bool(boggart.ConfigCameraHikVisionStreetEnabled) {
			device.Enable()
		} else {
			device.Disable()
		}

		c.devices.Register(boggart.DeviceCameraStreetID, device)
	}
}

func (c *Component) initPhones() {
	megafonPhone := c.config.String(boggart.ConfigMobileMegafonPhone)
	megafonPassword := c.config.String(boggart.ConfigMobileMegafonPassword)

	if megafonPhone == "" || megafonPassword == "" {
		return
	}

	providerMegafon := mobile.NewMegafon(megafonPhone, megafonPassword)
	device := devices.NewMegafonPhone(providerMegafon, c.config.Duration(boggart.ConfigMobileRepeatInterval))

	if c.config.Bool(boggart.ConfigMobileEnabled) {
		device.Enable()
	} else {
		device.Disable()
	}

	c.devices.Register(boggart.DevicePhoneID, device)
}
