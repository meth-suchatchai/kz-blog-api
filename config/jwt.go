package config

type JWT struct {
	Secret        string `mapstructure:"SECRET"`
	Issuer        string `mapstructure:"ISSUER"`
	Domain        string `mapstructure:"DOMAIN"`
	Expire        int64  `mapstructure:"EXPIRATION_TIME"`
	RefreshExpire int64  `mapstructure:"REFRESH_EXPIRATION_TIME"`
}
