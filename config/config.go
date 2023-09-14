package config

import (
	"github.com/spf13/viper"
	"strings"
)

var ENVIRONMENT = Env{}

type Env struct {
	Database `mapstructure:"DATABASE"`
	Redis    `mapstructure:"REDIS"`
	Server   `mapstructure:"SERVER"`
	TaxiMail `mapstructure:"TAXIMAIL"`
	JWT      `mapstructure:"JWT"`
	Storage  `mapstructure:"STORAGE"`
}

func ReadConfig(path string) (*Env, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(path)
	v.SetConfigType("toml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&ENVIRONMENT); err != nil {
		return nil, err
	}

	return &ENVIRONMENT, nil
}
