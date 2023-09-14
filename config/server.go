package config

type Server struct {
	Host            string `mapstructure:"HOST"`
	Port            int    `mapstructure:"PORT"`
	ApplicationName string `mapstructure:"APPLICATION_NAME"`
	Debug           bool   `mapstructure:"DEBUG"`
}
