package config

type ETCD struct {
	Hostname string `mapstructure:"hostname"`
	Timeout  int    `mapstructure:"timeout"`
}
