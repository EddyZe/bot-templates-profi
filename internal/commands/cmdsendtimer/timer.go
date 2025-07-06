package cmdsendtimer

import (
	"bot-templates-profi/internal/commands"
	"bot-templates-profi/internal/services/timerservice"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	barLen      = 10
	msgTemplate = "<b>Таймер</b>\n\n Осталось: <b>%s</b> \n\n<b>%s</b>"
)

type SendTimer[T commands.CommandType] struct {
	b  *bot.Bot
	ts timerservice.TimerService
}

func New[T commands.CommandType](b *bot.Bot, ts timerservice.TimerService) *SendTimer[T] {
	return &SendTimer[T]{
		b:  b,
		ts: ts,
	}
}

func (c *SendTimer[T]) Execute(ctx context.Context, arg T) {
	if v, ok := any(arg).(*models.Message); ok {
		c.handelMessage(ctx, v)
		return
	}
}

func (c *SendTimer[T]) handelMessage(ctx context.Context, msg *models.Message) {
	chatId := msg.Chat.ID
	splitText := strings.Split(msg.Text, " ")

	if len(splitText) != 2 {
		if _, err := c.b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatId,
				Text:   "Неверная команда. Введите /timer [кол-во секунд]. Например: /timer 30",
			}); err != nil {
			log.Println(err)
		}
		return
	}

	seconds, err := strconv.Atoi(splitText[1])
	if err != nil {
		if _, err := c.b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatId,
				Text:   "Кол-во секунд должно быть целым числом! Например: /timer 30",
			}); err != nil {
			log.Println(err)
		}
		return
	}

	total := seconds

	progresBar := c.ts.GenerateProgressBar(barLen, seconds, total)

	res, err := c.b.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID:    chatId,
			Text:      fmt.Sprintf(msgTemplate, c.formatDuration(seconds), progresBar),
			ParseMode: "html",
		})
	if err != nil {
		log.Println(err)
		return
	}

	go c.startTimer(ctx, res, seconds, total)
}

func (c *SendTimer[T]) startTimer(ctx context.Context, msg *models.Message, seconds, total int) {
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	defer timer.Stop()

	tick := 1
	ticker := time.NewTicker(time.Duration(tick) * time.Second)
	defer ticker.Stop()

	chatId := msg.Chat.ID

	for {
		select {
		case <-timer.C:
			if _, err := c.b.EditMessageText(
				ctx,
				&bot.EditMessageTextParams{
					ChatID:    chatId,
					Text:      "Время вышло! 🎉",
					MessageID: msg.ID,
				}); err != nil {
				log.Println(err)
			}
			return
		case <-ticker.C:
			seconds--
			if seconds > 0 {
				progresBar := c.ts.GenerateProgressBar(barLen, seconds, total)
				_, err := c.b.EditMessageText(
					ctx,
					&bot.EditMessageTextParams{
						ChatID:    chatId,
						Text:      fmt.Sprintf(msgTemplate, c.formatDuration(seconds), progresBar),
						ParseMode: "html",
						MessageID: msg.ID,
					})
				if err != nil {
					log.Println(err)
					return
				}
			}
		case <-ctx.Done():
			log.Println(ctx.Err())
			return
		}
	}
}

func (c *SendTimer[T]) formatDuration(seconds int) string {
	minutes := seconds / 60
	secs := seconds % 60
	if minutes > 0 {
		return fmt.Sprintf("%02dm %02ds", minutes, secs)
	}
	return fmt.Sprintf("%02ds", secs)
}
