package config

type Database struct {
	Host     string `mapstructure:"HOST"`
	Name     string `mapstructure:"NAME"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
	Port     int    `mapstructure:"PORT"`
	SSLMode  bool   `mapstructure:"SSLMODE"`
	Debug    bool   `mapstructure:"DEBUG"`
	Timezone string `mapstructure:"TIMEZONE"`
}
