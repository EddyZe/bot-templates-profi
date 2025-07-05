package config

import (
	"bot-templates-profi/pkg/configlouder"
)

var cfg AppConfig

type AppConfig struct {
	BotConfig *TelegramBotConfig
	Postgres  *PostgresConfig
}

type TelegramBotConfig struct {
	Token    string `env:"TELEGRAM_BOT_TOKEN"`
	Username string `env:"TELEGRAM_USERNAME"`
}

type PostgresConfig struct {
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	Name     string `env:"POSTGRES_NAME"`
}

func LoadConfig() *AppConfig {
	return configlouder.LoadEnvConfig(&cfg)
}
