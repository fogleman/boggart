package internal

import (
	"strconv"
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/protocols/rs485"
	"github.com/kihamo/boggart/components/boggart/providers/pulsar"
	"github.com/kihamo/shadow/components/config"
)

func (c *Component) ConfigVariables() []config.Variable {
	return []config.Variable{
		config.NewVariable(boggart.ConfigDevicesManagerCheckInterval, config.ValueTypeDuration).
			WithUsage("Health check interval").
			WithGroup("Device manager").
			WithEditable(true).
			WithDefault(time.Minute),
		config.NewVariable(boggart.ConfigDevicesManagerCheckTimeout, config.ValueTypeDuration).
			WithUsage("Health check timeout").
			WithGroup("Device manager").
			WithEditable(true).
			WithDefault(DefaultTimeoutChecker),
		config.NewVariable(boggart.ConfigListenerTelegramChats, config.ValueTypeString).
			WithUsage("Chats for messages").
			WithGroup("Listener Telegram").
			WithView([]string{config.ViewTags}).
			WithViewOptions(map[string]interface{}{config.ViewOptionTagsDefaultText: "add a chat ID"}),
		config.NewVariable(boggart.ConfigRS485Address, config.ValueTypeString).
			WithUsage("Serial port address").
			WithGroup("RS485 protocol").
			WithDefault(rs485.DefaultSerialAddress),
		config.NewVariable(boggart.ConfigRS485Timeout, config.ValueTypeDuration).
			WithUsage("Serial port timeout").
			WithGroup("RS485 protocol").
			WithDefault(rs485.DefaultTimeout),
		config.NewVariable(boggart.ConfigGPIOEnabled, config.ValueTypeBool).
			WithUsage("Enabled").
			WithGroup("GPIO"),
		config.NewVariable(boggart.ConfigGPIOPins, config.ValueTypeString).
			WithUsage("Pins listener").
			WithGroup("GPIO"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionHomeEnabled, config.ValueTypeBool).
			WithUsage("Enabled").
			WithGroup("Home video recorder"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionHomeRepeatInterval, config.ValueTypeDuration).
			WithUsage("Repeat interval").
			WithGroup("Home video recorder").
			WithDefault(time.Minute),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionHomeHost, config.ValueTypeString).
			WithUsage("Host").
			WithGroup("Home video recorder"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionHomePort, config.ValueTypeInt64).
			WithUsage("Port").
			WithGroup("Home video recorder"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionHomeUsername, config.ValueTypeString).
			WithUsage("Username").
			WithGroup("Home video recorder").
			WithDefault("admin"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionHomePassword, config.ValueTypeString).
			WithUsage("Password").
			WithGroup("Home video recorder").
			WithView([]string{config.ViewPassword}),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionVacationHomeEnabled, config.ValueTypeBool).
			WithUsage("Enabled").
			WithGroup("Vacation home video recorder"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionVacationHomeRepeatInterval, config.ValueTypeDuration).
			WithUsage("Repeat interval").
			WithGroup("Vacation home video recorder").
			WithDefault(time.Minute),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionVacationHomeHost, config.ValueTypeString).
			WithUsage("Host").
			WithGroup("Vacation home video recorder"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionVacationHomePort, config.ValueTypeInt64).
			WithUsage("Port").
			WithGroup("Vacation home video recorder"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionVacationHomeUsername, config.ValueTypeString).
			WithUsage("Username").
			WithGroup("Vacation home video recorder").
			WithDefault("admin"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionVacationHomePassword, config.ValueTypeString).
			WithUsage("Password").
			WithGroup("Vacation home video recorder").
			WithView([]string{config.ViewPassword}),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionGarageEnabled, config.ValueTypeBool).
			WithUsage("Enabled").
			WithGroup("Garage video recorder"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionGarageRepeatInterval, config.ValueTypeDuration).
			WithUsage("Repeat interval").
			WithGroup("Garage video recorder").
			WithDefault(time.Minute),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionGarageHost, config.ValueTypeString).
			WithUsage("Host").
			WithGroup("Garage video recorder"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionGaragePort, config.ValueTypeInt64).
			WithUsage("Port").
			WithGroup("Garage video recorder"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionGarageUsername, config.ValueTypeString).
			WithUsage("Username").
			WithGroup("Garage video recorder").
			WithDefault("admin"),
		config.NewVariable(boggart.ConfigVideoRecorderHikVisionGaragePassword, config.ValueTypeString).
			WithUsage("Password").
			WithGroup("Garage video recorder").
			WithView([]string{config.ViewPassword}),
		config.NewVariable(boggart.ConfigMercuryEnabled, config.ValueTypeBool).
			WithUsage("Enabled").
			WithGroup("Mercury devices"),
		config.NewVariable(boggart.ConfigMercuryRepeatInterval, config.ValueTypeDuration).
			WithUsage("Repeat interval").
			WithGroup("Mercury devices").
			WithDefault(time.Minute * 2),
		config.NewVariable(boggart.ConfigMercuryDeviceAddress, config.ValueTypeString).
			WithUsage("Device address in format XXXXXX (last 6 digits of device serial number)").
			WithGroup("Mercury devices"),
		config.NewVariable(boggart.ConfigMikrotikEnabled, config.ValueTypeBool).
			WithUsage("Enabled").
			WithGroup("Mikrotik devices"),
		config.NewVariable(boggart.ConfigMikrotikRepeatInterval, config.ValueTypeDuration).
			WithUsage("Repeat interval").
			WithGroup("Mikrotik devices").
			WithDefault(time.Minute * 5),
		config.NewVariable(boggart.ConfigMikrotikAddress, config.ValueTypeString).
			WithUsage("API address in format host:port").
			WithGroup("Mikrotik devices").
			WithDefault("192.168.88.1:8728"),
		config.NewVariable(boggart.ConfigMikrotikUsername, config.ValueTypeString).
			WithUsage("Username").
			WithGroup("Mikrotik devices").
			WithDefault("admin"),
		config.NewVariable(boggart.ConfigMikrotikPassword, config.ValueTypeString).
			WithUsage("Password").
			WithGroup("Mikrotik devices").
			WithView([]string{config.ViewPassword}),
		config.NewVariable(boggart.ConfigMikrotikTimeout, config.ValueTypeDuration).
			WithUsage("Request timeout").
			WithGroup("Mikrotik devices").
			WithDefault(time.Second * 10),
		config.NewVariable(boggart.ConfigMobileEnabled, config.ValueTypeBool).
			WithUsage("Enabled").
			WithGroup("Mobile accounts"),
		config.NewVariable(boggart.ConfigMobileRepeatInterval, config.ValueTypeDuration).
			WithUsage("Repeat interval").
			WithGroup("Mobile accounts").
			WithDefault(time.Minute * 30),
		config.NewVariable(boggart.ConfigMobileMegafonPhone, config.ValueTypeString).
			WithUsage("Phone number").
			WithGroup("Mobile Megafon accounts"),
		config.NewVariable(boggart.ConfigMobileMegafonPassword, config.ValueTypeString).
			WithUsage("Password").
			WithGroup("Mobile Megafon accounts").
			WithView([]string{config.ViewPassword}),
		config.NewVariable(boggart.ConfigPulsarEnabled, config.ValueTypeBool).
			WithUsage("Enabled").
			WithGroup("Pulsar devices"),
		config.NewVariable(boggart.ConfigPulsarRepeatInterval, config.ValueTypeDuration).
			WithUsage("Repeat interval").
			WithGroup("Pulsar devices").
			WithDefault(time.Minute * 15),
		config.NewVariable(boggart.ConfigPulsarHeatMeterAddress, config.ValueTypeString).
			WithUsage("Device address HEX value (AABBCCDD). If empty system try to find device").
			WithGroup("Pulsar devices"),
		config.NewVariable(boggart.ConfigPulsarColdWaterSerialNumber, config.ValueTypeString).
			WithUsage("Device address of cold water meter in format AABBCCDD").
			WithGroup("Pulsar devices"),
		config.NewVariable(boggart.ConfigPulsarColdWaterPulseInput, config.ValueTypeUint64).
			WithUsage("Input number of cold water meter in heat meter").
			WithGroup("Pulsar devices").
			WithDefault(pulsar.Input1).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{strconv.FormatUint(pulsar.Input1, 10), "#1"},
					{strconv.FormatUint(pulsar.Input2, 10), "#2"},
				},
			}),
		config.NewVariable(boggart.ConfigPulsarColdWaterStartValue, config.ValueTypeFloat64).
			WithUsage("Start value of cold water meter (in m3)").
			WithGroup("Pulsar devices").
			WithDefault(0),
		config.NewVariable(boggart.ConfigPulsarHotWaterSerialNumber, config.ValueTypeString).
			WithUsage("Device address of hot water meter in format AABBCCDD").
			WithGroup("Pulsar devices"),
		config.NewVariable(boggart.ConfigPulsarHotWaterPulseInput, config.ValueTypeUint64).
			WithUsage("Input number of hot water meter").
			WithGroup("Pulsar devices").
			WithDefault(pulsar.Input2).
			WithView([]string{config.ViewEnum}).
			WithViewOptions(map[string]interface{}{
				config.ViewOptionEnumOptions: [][]interface{}{
					{strconv.FormatUint(pulsar.Input1, 10), "#1"},
					{strconv.FormatUint(pulsar.Input2, 10), "#2"},
				},
			}),
		config.NewVariable(boggart.ConfigPulsarHotWaterStartValue, config.ValueTypeFloat64).
			WithUsage("Start value of hot water (in m3)").
			WithGroup("Pulsar devices").
			WithDefault(0),
		config.NewVariable(boggart.ConfigSoftVideoEnabled, config.ValueTypeBool).
			WithUsage("Enabled").
			WithGroup("SoftVideo provider"),
		config.NewVariable(boggart.ConfigSoftVideoRepeatInterval, config.ValueTypeDuration).
			WithUsage("Repeat interval").
			WithGroup("SoftVideo provider").
			WithDefault(time.Hour * 12),
		config.NewVariable(boggart.ConfigSoftVideoLogin, config.ValueTypeString).
			WithUsage("Login").
			WithGroup("SoftVideo provider"),
		config.NewVariable(boggart.ConfigSoftVideoPassword, config.ValueTypeString).
			WithUsage("Password").
			WithGroup("SoftVideo provider").
			WithView([]string{config.ViewPassword}),
		config.NewVariable(boggart.ConfigSensorBME280Enabled, config.ValueTypeBool).
			WithUsage("Enabled").
			WithGroup("Sensor 280"),
		config.NewVariable(boggart.ConfigSensorBME280RepeatInterval, config.ValueTypeDuration).
			WithUsage("Repeat interval").
			WithGroup("Sensor 280").
			WithDefault(time.Minute),
		config.NewVariable(boggart.ConfigSensorBME280Bus, config.ValueTypeInt).
			WithUsage("Bus number").
			WithGroup("Sensor 280").
			WithDefault(1),
		config.NewVariable(boggart.ConfigSensorBME280Address, config.ValueTypeInt).
			WithUsage("Address").
			WithGroup("Sensor 280").
			WithDefault(0x76),
		config.NewVariable(boggart.ConfigOwnTracksEnabled, config.ValueTypeBool).
			WithUsage("Enabled").
			WithGroup("OwnTracks"),
	}
}

func (c *Component) ConfigWatchers() []config.Watcher {
	return []config.Watcher{
		config.NewWatcher([]string{
			boggart.ConfigDevicesManagerCheckInterval,
		}, c.watchDevicesManagerCheckInterval),
		config.NewWatcher([]string{
			boggart.ConfigDevicesManagerCheckTimeout,
		}, c.watchDevicesManagerCheckTimeout),
	}
}

func (c *Component) watchDevicesManagerCheckInterval(_ string, newValue interface{}, _ interface{}) {
	c.devicesManager.SetCheckerTickerDuration(newValue.(time.Duration))
}

func (c *Component) watchDevicesManagerCheckTimeout(_ string, newValue interface{}, _ interface{}) {
	c.devicesManager.SetCheckerTimeout(newValue.(time.Duration))
}
