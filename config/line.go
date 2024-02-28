package config

type Line struct {
	LineApi     string `mapstructure:"LINE_API"`
	BotApi      string `mapstructure:"BOT_API"`
	AccessToken string `mapstructure:"ACCESS_TOKEN"`
}
