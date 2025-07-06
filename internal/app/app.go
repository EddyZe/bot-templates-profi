package app

import (
	"bot-templates-profi/internal/app/bot"
	"bot-templates-profi/internal/app/db/psqldb"
	"bot-templates-profi/internal/config"
	"bot-templates-profi/internal/handlers/telegramhandl"
	"bot-templates-profi/internal/repositories/userrepo"
	"bot-templates-profi/internal/services/ieservice"
	"bot-templates-profi/internal/services/timerservice"
	"bot-templates-profi/internal/services/userservice"
)

func Run(cfg *config.AppConfig) error {
	psql := psqldb.MustRun(cfg.Postgres)

	ur := userrepo.New(psql)

	us := userservice.New(ur)
	ies := ieservice.New()
	ts := timerservice.New("🟩", "⬜")

	handler := telegramhandl.New(us, ies, ts)

	tgBot, err := bot.New(cfg.BotConfig, handler)
	if err != nil {
		return err
	}

	if err := tgBot.Start(); err != nil {
		return err
	}

	return nil
}

func MustRun(cfg *config.AppConfig) {
	if err := Run(cfg); err != nil {
		panic(err)
	}
}
