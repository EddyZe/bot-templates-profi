package cmdrandloc

import (
	"bot-templates-profi/internal/commands"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"math/rand"
	"time"
)

type RandLocation[T commands.CommandType] struct {
	b *bot.Bot
}

func New[T commands.CommandType](b *bot.Bot) *RandLocation[T] {
	return &RandLocation[T]{
		b: b,
	}
}

func (t *RandLocation[T]) Execute(ctx context.Context, arg T) {
	if v, ok := any(arg).(*models.Message); ok {
		t.handleMessage(ctx, t.b, v)
		return
	}
}

func (t *RandLocation[T]) handleMessage(ctx context.Context, b *bot.Bot, msg *models.Message) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	x := r.Float64() * 90
	y := r.Float64() * 90

	log.Println(x)
	log.Println(y)

	chatId := msg.Chat.ID

	if _, err := b.SendLocation(
		ctx,
		&bot.SendLocationParams{
			ChatID:    chatId,
			Longitude: x,
			Latitude:  y,
		},
	); err != nil {
		log.Println("Error sending location: ", err)
	}

}
