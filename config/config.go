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
	JWT      `mapstructure:"JWT"`
	Storage  `mapstructure:"STORAGE"`
	ETCD     `mapstructure:"ETCD"`
	//Line    `mapstructure:"LINE"`
	//Kafka    `mapstructure:"KAFKA"`
	//TaxiMail `mapstructure:"TAXIMAIL"`
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
