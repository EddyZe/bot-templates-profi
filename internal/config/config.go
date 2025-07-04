package config

import "bot-templates-profi/internal/pkg/configlouder"

var cfg AppConfig

type AppConfig struct {
	BotConfig *TelegramBotConfig
}

type TelegramBotConfig struct {
	Token    string `env:"TELEGRAM_BOT_TOKEN"`
	Username string `env:"TELEGRAM_USERNAME"`
}

func LoadConfig() *AppConfig {
	return configlouder.LoadEnvConfig(&cfg)
}
