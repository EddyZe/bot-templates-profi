package commands

import (
	"context"
	"github.com/go-telegram/bot/models"
)

const (
	Start        = "/start"
	RandLocation = "/rand"
	MyLocation   = "/mylocation"
	Timer        = "/timer"
	All          = "/all"
)

type CommandType interface {
	*models.Message | *models.CallbackQuery
}

type Command[T CommandType] interface {
	Execute(ctx context.Context, arg T)
}
