package cmdall

import (
	"bot-templates-profi/internal/commands"
	"bot-templates-profi/internal/services/ieservice"
	"bot-templates-profi/internal/services/userservice"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"os"
	"time"
)

type AllUsers[T commands.CommandType] struct {
	b   *bot.Bot
	us  userservice.UserService
	ies ieservice.IEService
}

func New[T commands.CommandType](b *bot.Bot, us userservice.UserService, ies ieservice.IEService) *AllUsers[T] {
	return &AllUsers[T]{
		b:   b,
		us:  us,
		ies: ies,
	}
}

func (s *AllUsers[T]) Execute(ctx context.Context, arg T) {
	if v, ok := any(arg).(*models.Message); ok {
		chatId := v.Chat.ID
		usrs := s.us.FindAll(ctx)

		f, err := s.ies.Export(usrs)
		if err != nil {
			log.Println(err)
			if _, err := s.b.SendMessage(
				ctx,
				&bot.SendMessageParams{
					ChatID: chatId,
					Text:   "Не удалось записать csv файл",
				}); err != nil {
				log.Println(err)
			}
			return
		}

		openFile, err := os.Open(f.Name())
		defer func(openFile *os.File) {
			err := openFile.Close()
			if err != nil {
				log.Println(err)
			}
		}(openFile)

		if _, err := s.b.SendDocument(
			ctx,
			&bot.SendDocumentParams{
				ChatID: chatId,
				Caption: fmt.Sprintf(
					"Список пользователей на %s",
					time.Now().Format("02-01-2006 15:04"),
				),
				Document: &models.InputFileUpload{
					Filename: "usr.csv",
					Data:     openFile,
				},
			}); err != nil {
			log.Println(err)
		}
	}
}
