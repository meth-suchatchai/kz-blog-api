package config

type TaxiMail struct {
	ApiKey      string `mapstructure:"API_KEY"`
	SecretKey   string `mapstructure:"SECRET_KEY"`
	URL         string `mapstructure:"URL"`
	SMSTemplate string `mapstructure:"SMS_TEMPLATE"`
}
