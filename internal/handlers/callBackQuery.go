package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
)

type CallbackQuery struct {
	Prefix   string
	LevelReq AccessLevel
	ChatType ChatType
	Handler  func(*bot.Server, *gotgbot.Bot, *ext.Context) error
}

func GetCallbackQueryList() []CallbackQuery {
	CallbackQueries := make([]CallbackQuery, 0)

	return CallbackQueries
}
