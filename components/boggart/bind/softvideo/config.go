package softvideo

import (
	"time"
)

const (
	DefaultUpdaterInterval = time.Hour
)

type Config struct {
	Login           string        `valid:"required"`
	Password        string        `valid:"required"`
	UpdaterInterval time.Duration `mapstructure:"updater_interval" yaml:"updater_interval"`
}

func (Type) Config() interface{} {
	return &Config{
		UpdaterInterval: DefaultUpdaterInterval,
	}
}