package config

type ETCD struct {
	Hostname []string `mapstructure:"hostname"`
	Timeout  int      `mapstructure:"timeout"`
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`
}
