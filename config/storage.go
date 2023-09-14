package config

type Storage struct {
	Bucket      string `mapstructure:"BUCKET"`
	Endpoint    string `mapstructure:"ENDPOINT"`
	AccessKeyId string `mapstructure:"ACCESS_KEY_ID"`
	SecretKey   string `mapstructure:"SECRET_KEY"`
	EnableSSL   bool   `mapstructure:"ENABLE_SSL"`
	Region      string `mapstructure:"REGION"`
}
