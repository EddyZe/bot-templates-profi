package cmdstart

import (
	"bot-templates-profi/internal/commands"
	"bot-templates-profi/internal/domain/entity"
	"bot-templates-profi/internal/repositories/userrepo"
	"bot-templates-profi/internal/services/userservice"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

type CmdStart[T commands.CommandType] struct {
	b  *bot.Bot
	us userservice.UserService
}

func New[T commands.CommandType](b *bot.Bot, us userservice.UserService) *CmdStart[T] {
	return &CmdStart[T]{
		b:  b,
		us: us,
	}
}

func (c *CmdStart[T]) Execute(ctx context.Context, arg T) {
	if v, ok := any(arg).(*models.Message); ok {
		c.executeMessage(ctx, v)
		return
	}
}

func (c *CmdStart[T]) executeMessage(ctx context.Context, v *models.Message) {
	chatId := v.Chat.ID
	userName := v.Chat.Username

	newUser := entity.User{
		TelegramId: v.Chat.ID,
		Username:   userName,
	}

	if err := c.us.CreateUser(ctx, &newUser); err != nil {
		if err.Error() == userrepo.UserIsExist {
			if err := c.us.UpdateByTelegramId(ctx, &newUser); err != nil {
				if err.Error() == userrepo.UserNotFound {
					log.Println(err)
					return
				} else {
					if _, err := c.b.SendMessage(
						ctx,
						&bot.SendMessageParams{
							ChatID: chatId,
							Text:   "Ошибка на стороне сервера.",
						}); err != nil {
						log.Println(err)
						return
					}
				}
			}
		} else {
			if _, err := c.b.SendMessage(
				ctx,
				&bot.SendMessageParams{
					ChatID: chatId,
					Text:   "Ошибка на стороне сервера.",
				}); err != nil {
				log.Println(err)
				return
			}
		}
	}

	_, err := c.b.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID: chatId,
			Text:   "Привет.",
		})
	if err != nil {
		log.Println("Error sending message:", err)
		return
	}

}
