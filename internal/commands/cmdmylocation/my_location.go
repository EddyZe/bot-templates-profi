package cmdmylocation

import (
	"bot-templates-profi/internal/commands"
	"bot-templates-profi/internal/util/inlinebtns"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

type MyLocation[T commands.CommandType] struct {
	b *bot.Bot
}

func New[T commands.CommandType](b *bot.Bot) *MyLocation[T] {
	return &MyLocation[T]{b: b}
}

func (t *MyLocation[T]) Execute(ctx context.Context, arg T) {
	if v, ok := any(arg).(*models.Message); ok {
		if v.Location != nil {
			t.handleLocation(ctx, v)
			return
		}
		t.handleMessage(ctx, v)
		return
	}
}

func (t *MyLocation[T]) handleMessage(ctx context.Context, msg *models.Message) {
	chatId := msg.Chat.ID

	if _, err := t.b.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID: chatId,
			Text:   "Поделитесь геолокацией",
			ReplyMarkup: &models.ReplyKeyboardMarkup{
				ResizeKeyboard:  true,
				OneTimeKeyboard: true,
				Keyboard: [][]models.KeyboardButton{
					{
						models.KeyboardButton{
							Text:            inlinebtns.SendUserLocationText,
							RequestLocation: true,
						},
					},
				},
			},
		}); err != nil {
		log.Println(err)
	}
}

func (t *MyLocation[T]) handleLocation(ctx context.Context, msg *models.Message) {
	chatId := msg.Chat.ID
	location := msg.Location

	if _, err := t.b.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID: chatId,
			Text:   "Ваша геопозиция",
		}); err != nil {
		log.Println(err)
	}

	if _, err := t.b.SendLocation(
		ctx,
		&bot.SendLocationParams{
			ChatID:    chatId,
			Longitude: location.Longitude,
			Latitude:  location.Latitude,
		},
	); err != nil {
		log.Println("Error sending location: ", err)
	}
}
