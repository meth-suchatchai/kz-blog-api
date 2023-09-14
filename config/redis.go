package config

type Redis struct {
	Address  string `mapstructure:"ADDRESS"`
	Password string `mapstructure:"PASSWORD"`
	Channel  int    `mapstructure:"CHANNEL"`
	Expire   int    `mapstructure:"EXPIRE"`
	Timeout  int    `mapstructure:"TIMEOUT"`
}
