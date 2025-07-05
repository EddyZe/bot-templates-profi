package cmdsendtimer

import (
	"bot-templates-profi/internal/commands"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"time"
)

type SendTimer[T commands.CommandType] struct {
	b *bot.Bot
}

func New[T commands.CommandType](b *bot.Bot) *SendTimer[T] {
	return &SendTimer[T]{b: b}
}

func (c *SendTimer[T]) Execute(ctx context.Context, arg T) {
	if v, ok := any(arg).(*models.Message); ok {
		c.handelMessage(ctx, v)
		return
	}
}

func (c *SendTimer[T]) handelMessage(ctx context.Context, msg *models.Message) {
	chatId := msg.Chat.ID

	//f, err := os.Open("videos/timer10.mp4")
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	res, err := c.b.SendVideo(
		ctx,
		&bot.SendVideoParams{
			ChatID: chatId,
			Video: &models.InputFileString{
				Data: "BAACAgIAAxkDAAIWh2hpRHwQ3-Ky8qYAAVFphXCNzmKZrwACtnsAAvFzSUthdC8FRgOATDYE",
			},
		})

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res.Video.FileID)

	go func() {
		timer := time.NewTimer(10 * time.Second)
		defer timer.Stop()

		<-timer.C
		if _, err := c.b.DeleteMessage(
			ctx,
			&bot.DeleteMessageParams{
				MessageID: res.ID,
				ChatID:    chatId,
			},
		); err != nil {
			log.Println(err)
		}

		if _, err := c.b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatId,
				Text:   "Время вышло!",
			}); err != nil {
			log.Println(err)
			return
		}
	}()
}
