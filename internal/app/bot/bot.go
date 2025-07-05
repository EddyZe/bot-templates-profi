package bot

import (
	"bot-templates-profi/internal/config"
	"bot-templates-profi/internal/handlers/telegramhandl"
	"context"
	"errors"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
	"syscall"
)

type TgBot struct {
	cfg     *config.TelegramBotConfig
	handler telegramhandl.Handler
}

func New(cfg *config.TelegramBotConfig, handler telegramhandl.Handler) (*TgBot, error) {
	if cfg.Token == "" {
		return nil, errors.New("telegram bot token is required")
	}

	return &TgBot{
		cfg:     cfg,
		handler: handler,
	}, nil
}

func (t *TgBot) Start() error {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	opts := []bot.Option{
		bot.WithDefaultHandler(t.handler.Handle),
	}

	b, err := bot.New(t.cfg.Token, opts...)
	if err != nil {
		return err
	}

	b.Start(ctx)

	return nil
}
