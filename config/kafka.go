package config

type Kafka struct {
	Hostname string `mapstructure:"HOSTNAME"`
	Acks     string `mapstructure:"ACKS"`
}
