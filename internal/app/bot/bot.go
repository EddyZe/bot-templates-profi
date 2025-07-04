package bot

import (
	"bot-templates-profi/internal/config"
	"bot-templates-profi/internal/handlers/bothandler"
	"context"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
)

type TgBot struct {
	*bot.Bot
}

func New(cfg *config.TelegramBotConfig) (*TgBot, error) {
	opts := []bot.Option{
		bot.WithDefaultHandler(bothandler.DefaultHandler),
	}

	b, err := bot.New(cfg.Token, opts...)
	if err != nil {
		return nil, err
	}

	return &TgBot{
		b,
	}, nil
}

func (t *TgBot) Start() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	t.Bot.Start(ctx)
}
