package bothandler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func DefaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		_, _ = b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   update.Message.Text,
			},
		)
	}
}
