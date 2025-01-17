package timelapse

import (
	"os"
	"time"

	"github.com/kihamo/boggart/components/boggart"
	"github.com/kihamo/boggart/components/boggart/di"
	"github.com/kihamo/boggart/components/mqtt"
	"github.com/kihamo/boggart/types"
)

type Config struct {
	di.ProbesConfig `mapstructure:",squash" yaml:",inline"`
	di.LoggerConfig `mapstructure:",squash" yaml:",inline"`

	UpdaterInterval   time.Duration  `mapstructure:"updater_interval" yaml:"updater_interval"`
	SaveDirectory     string         `valid:"required" mapstructure:"save_directory" yaml:"save_directory"`
	CaptureURL        types.URL      `valid:"required" mapstructure:"capture_url" yaml:"capture_url"`
	FileNameFormat    string         `mapstructure:"file_name_format" yaml:"file_name_format"`
	FileMode          types.FileMode `mapstructure:"file_mode" yaml:"file_mode"`
	SaveDirectoryMode types.FileMode `mapstructure:"save_directory_mode" yaml:"save_directory_mode"`
	FilesOnPage       int            `mapstructure:"files_on_page" yaml:"files_on_page"`
	TopicCapture      mqtt.Topic     `mapstructure:"topic_capture" yaml:"topic_capture"`
}

func (Type) ConfigDefaults() interface{} {
	probesConfig := di.ProbesConfigDefaults()
	probesConfig.ReadinessPeriod = time.Minute * 10

	return &Config{
		ProbesConfig:    probesConfig,
		LoggerConfig:    di.LoggerConfigDefaults(),
		UpdaterInterval: time.Hour,
		FileNameFormat:  "20060102_150405",
		FileMode: types.FileMode{
			FileMode: os.FileMode(0664),
		},
		SaveDirectoryMode: types.FileMode{
			FileMode: os.FileMode(0774),
		},
		FilesOnPage:  9,
		TopicCapture: boggart.ComponentName + "/timelapse/+/capture",
	}
}
