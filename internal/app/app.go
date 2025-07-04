package app

import (
	"bot-templates-profi/internal/app/bot"
	"bot-templates-profi/internal/config"
)

func Run(cfg *config.AppConfig) error {

	tgBot, err := bot.New(cfg.BotConfig)
	if err != nil {
		return err
	}
	tgBot.Start()

	return nil
}

func MustRun(cfg *config.AppConfig) {
	if err := Run(cfg); err != nil {
		panic(err)
	}
}
