package hikvision

import (
	"context"
	"encoding/xml"
	"errors"
	"net/http"
	"strconv"
	"time"

	tracing "github.com/kihamo/shadow/components/tracing/http"
)

const (
	proxyPTZPrefixURL = "/PTZCtrl"
)

type PTZStatus struct {
	AbsoluteHigh PTZDataAbsoluteHigh `xml:"AbsoluteHigh"`
}

type PTZChannelList struct {
	Channels []PTZChannel `xml:"PTZChannel"`
}

type PTZChannel struct {
	ID                   uint64 `xml:"id"`
	Enabled              bool   `xml:"enabled"`
	VideoInputID         uint64 `xml:"videoInputID"`
	PanMaxSpeed          uint64 `xml:"panMaxSpeed"`
	TiltMaxSpeed         uint64 `xml:"tiltMaxSpeed"`
	PresetSpeed          uint64 `xml:"presetSpeed"`
	AutoPatrolSpeed      uint64 `xml:"autoPatrolSpeed"`
	KeyBoardControlSpeed string `xml:"keyBoardControlSpeed"`
	ControlProtocol      string `xml:"controlProtocol"`
	DefaultPresetID      uint64 `xml:"defaultPresetID"`
	PanSupport           bool   `xml:"panSupport"`
	TiltSupport          bool   `xml:"tiltSupport"`
	ZoomSupport          bool   `xml:"zoomSupport"`
	ManualControlSpeed   string `xml:"manualControlSpeed"`
}

type PTZData struct {
	XMLName      xml.Name             `xml:"PTZData"`
	Relative     *PTZDataRelative     `xml:"Relative,omitempty"`
	AbsoluteHigh *PTZDataAbsoluteHigh `xml:"AbsoluteHigh,omitempty"`
}

type PTZDataContinuous struct {
	XMLName   xml.Name `xml:"PTZData"`
	Pan       int64    `xml:"pan,omitempty"`
	Tilt      int64    `xml:"tilt,omitempty"`
	Zoom      int64    `xml:"zoom,omitempty"`
	Momentary struct {
		Duration int64 `xml:"duration"`
	} `xml:"Momentary,omitempty"`
}

type PTZDataRelative struct {
	PositionX    int64 `xml:"positionX,omitempty"`
	PositionY    int64 `xml:"positionY,omitempty"`
	RelativeZoom int64 `xml:"relativeZoom,omitempty"`
}

type PTZDataAbsoluteHigh struct {
	Elevation    int64  `xml:"elevation,omitempty"`
	Azimuth      uint64 `xml:"azimuth,omitempty"`
	AbsoluteZoom uint64 `xml:"absoluteZoom,omitempty"`
}

func (a *ISAPI) PTZChannels(ctx context.Context) (list PTZChannelList, err error) {
	u := a.address + proxyPTZPrefixURL + "/channels"

	ctx = tracing.OperationNameToContext(ctx, ComponentName+".PTZChannels")

	err = a.DoXML(ctx, http.MethodGet, u, nil, &list)
	return list, err
}

func (a *ISAPI) PTZStatus(ctx context.Context, channel uint64) (status PTZStatus, err error) {
	u := a.address + proxyPTZPrefixURL + "/channels/" + strconv.FormatUint(channel, 10) + "/status"

	ctx = tracing.OperationNameToContext(ctx, ComponentName+".PTZStatus")

	err = a.DoXML(ctx, http.MethodGet, u, nil, &status)
	return status, err
}

func (a *ISAPI) PTZPresetGoTo(ctx context.Context, channel, preset uint64) error {
	u := a.address + proxyPTZPrefixURL + "/channels/" +
		strconv.FormatUint(channel, 10) +
		"/presets/" +
		strconv.FormatUint(preset, 10) +
		"/goto"

	result := ResponseStatus{}

	ctx = tracing.OperationNameToContext(ctx, ComponentName+".PTZPresetGoTo")

	err := a.DoXML(ctx, http.MethodPut, u, nil, &result)
	if err != nil {
		return err
	}

	if result.StatusCode != 1 {
		return errors.New(result.StatusString)
	}

	return nil
}

func (a *ISAPI) PTZRelative(ctx context.Context, channel uint64, x, y, zoom int64) error {
	if zoom < -100 {
		zoom = -100
	} else if zoom > 100 {
		zoom = 100
	}

	u := a.address + proxyPTZPrefixURL + "/channels/" + strconv.FormatUint(channel, 10) + "/relative"

	data := PTZData{
		Relative: &PTZDataRelative{
			PositionX:    x,
			PositionY:    y,
			RelativeZoom: zoom,
		},
	}

	result := ResponseStatus{}

	ctx = tracing.OperationNameToContext(ctx, ComponentName+".PTZRelative")

	err := a.DoXML(ctx, http.MethodPut, u, data, &result)
	if err != nil {
		return err
	}

	if result.StatusCode != 1 {
		return errors.New(result.StatusString)
	}

	return nil
}

func (a *ISAPI) PTZAbsolute(ctx context.Context, channel uint64, elevation int64, azimuth, absoluteZoom uint64) error {
	if elevation < -900 {
		elevation = -900
	} else if elevation > 2700 {
		elevation = 2700
	}

	if azimuth < 0 {
		azimuth = 0
	} else if azimuth > 3600 {
		azimuth = 3600
	}

	if absoluteZoom < 0 {
		absoluteZoom = 0
	} else if absoluteZoom > 1000 {
		absoluteZoom = 1000
	}

	u := a.address + proxyPTZPrefixURL + "/channels/" + strconv.FormatUint(channel, 10) + "/absolute"

	data := PTZData{
		AbsoluteHigh: &PTZDataAbsoluteHigh{
			Elevation:    elevation,
			Azimuth:      azimuth,
			AbsoluteZoom: absoluteZoom,
		},
	}

	result := ResponseStatus{}

	ctx = tracing.OperationNameToContext(ctx, ComponentName+".PTZAbsolute")

	err := a.DoXML(ctx, http.MethodPut, u, data, &result)
	if err != nil {
		return err
	}

	if result.StatusCode != 1 {
		return errors.New(result.StatusString)
	}

	return nil
}

func (a *ISAPI) PTZContinuous(ctx context.Context, channel uint64, pan, tilt, zoom int64) error {
	if pan < -100 {
		pan = -100
	} else if pan > 100 {
		pan = 100
	}

	if tilt < -100 {
		tilt = -100
	} else if tilt > 100 {
		tilt = 100
	}

	if zoom < -100 {
		zoom = -100
	} else if zoom > 100 {
		zoom = 100
	}

	u := a.address + proxyPTZPrefixURL + "/channels/" + strconv.FormatUint(channel, 10) + "/continuous"

	data := PTZDataContinuous{
		Pan:  pan,
		Tilt: tilt,
		Zoom: zoom,
	}

	result := ResponseStatus{}

	ctx = tracing.OperationNameToContext(ctx, ComponentName+".PTZContinuous")

	err := a.DoXML(ctx, http.MethodPut, u, data, &result)
	if err != nil {
		return err
	}

	if result.StatusCode != 1 {
		return errors.New(result.StatusString)
	}

	return nil
}

func (a *ISAPI) PTZMomentary(ctx context.Context, channel uint64, pan, tilt, zoom int64, duration time.Duration) error {
	if pan < -100 {
		pan = -100
	} else if pan > 100 {
		pan = 100
	}

	if tilt < -100 {
		tilt = -100
	} else if tilt > 100 {
		tilt = 100
	}

	if zoom < -100 {
		zoom = -100
	} else if zoom > 100 {
		zoom = 100
	}

	if duration < 0 {
		duration = 0
	}

	u := a.address + proxyPTZPrefixURL + "/channels/" + strconv.FormatUint(channel, 10) + "/momentary"

	data := PTZDataContinuous{
		Pan:  pan,
		Tilt: tilt,
		Zoom: zoom,
	}
	data.Momentary.Duration = int64(duration.Seconds() * 1000)

	result := ResponseStatus{}

	ctx = tracing.OperationNameToContext(ctx, ComponentName+".PTZMomentary")

	err := a.DoXML(ctx, http.MethodPut, u, data, &result)
	if err != nil {
		return err
	}

	if result.StatusCode != 1 {
		return errors.New(result.StatusString)
	}

	return nil
}
