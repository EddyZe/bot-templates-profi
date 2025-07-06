package telegramhandl

import (
	"bot-templates-profi/internal/commands"
	"bot-templates-profi/internal/commands/cmdall"
	"bot-templates-profi/internal/commands/cmdmylocation"
	"bot-templates-profi/internal/commands/cmdrandloc"
	"bot-templates-profi/internal/commands/cmdsendtimer"
	"bot-templates-profi/internal/commands/cmdstart"
	"bot-templates-profi/internal/services/ieservice"
	"bot-templates-profi/internal/services/timerservice"
	"bot-templates-profi/internal/services/userservice"
	"context"
	"errors"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"strings"
)

const (
	CommandNotFound    = "command not found"
	UnknownTypeMessage = "unknown type of message"
)

type TelegramHandler struct {
	us  userservice.UserService
	ies ieservice.IEService
	ts  timerservice.TimerService
}

func New(us userservice.UserService, ies ieservice.IEService, ts timerservice.TimerService) *TelegramHandler {
	return &TelegramHandler{
		us:  us,
		ies: ies,
		ts:  ts,
	}
}

func (t *TelegramHandler) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		go func() {
			cmd, err := t.handleMessage(b, update.Message)
			if err != nil {
				t.ErrResponse(ctx, b, update.Message.Chat.ID, err)
				return
			}
			cmd.Execute(ctx, update.Message)
		}()
		return
	}

	if update.EditedMessage != nil {
		go func() {
			cmd, err := t.handleEditMessage(b, update.EditedMessage)
			if err != nil {
				t.ErrResponse(ctx, b, update.EditedMessage.Chat.ID, err)
				return
			}
			cmd.Execute(ctx, update.EditedMessage)
		}()
		return
	}
}

func (t *TelegramHandler) handleMessage(b *bot.Bot, message *models.Message) (commands.Command[*models.Message], error) {
	if message.Text != "" {
		if message.Entities != nil {
			entity := message.Entities
			for _, e := range entity {
				if e.Type == "custom_emoji" {
					log.Println(e.CustomEmojiID)
				}
			}
		}
		return t.handelTextCommand(b, message)
	}

	if message.Location != nil {
		return cmdmylocation.New[*models.Message](b), nil
	}

	return nil, errors.New(UnknownTypeMessage)
}

func (t *TelegramHandler) handleEditMessage(b *bot.Bot, message *models.Message) (commands.Command[*models.Message], error) {
	if message.Location != nil {
		return cmdmylocation.New[*models.Message](b), nil
	}

	return nil, errors.New(UnknownTypeMessage)
}

func (t *TelegramHandler) handelTextCommand(b *bot.Bot, message *models.Message) (commands.Command[*models.Message], error) {
	text := message.Text

	if text == commands.Start {
		cmd := cmdstart.New[*models.Message](b, t.us)
		return cmd, nil
	}

	if text == commands.RandLocation {
		cmd := cmdrandloc.New[*models.Message](b)
		return cmd, nil
	}

	if text == commands.MyLocation {
		cmd := cmdmylocation.New[*models.Message](b)
		return cmd, nil
	}

	if strings.HasPrefix(text, commands.Timer) {
		cmd := cmdsendtimer.New[*models.Message](b, t.ts)
		return cmd, nil
	}

	if text == commands.All {
		cmd := cmdall.New[*models.Message](b, t.us, t.ies)
		return cmd, nil
	}

	return nil, errors.New(CommandNotFound)
}

func (t *TelegramHandler) ErrResponse(ctx context.Context, b *bot.Bot, chatId int64, err error) {
	switch err.Error() {
	case UnknownTypeMessage:
		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatId,
				Text:   "😔 Пока-что я не умею обрабатывать такие сообщения.",
			},
		); err != nil {
			log.Println("Error of sending message on error: ", err)
			return
		}
		break
	case CommandNotFound:
		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatId,
				Text:   "😔 Пока-что я не знаю такой кооманды!",
			},
		); err != nil {
			log.Println("Error of sending message on error: ", err)
			return
		}
		break
	}
}
