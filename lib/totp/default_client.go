package totp

type Config struct {
	AppName string
}

type defaultClient struct {
	name string
}

func New(config Config) Client {
	return &defaultClient{name: config.AppName}
}
