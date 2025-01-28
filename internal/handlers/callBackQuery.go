package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handlers/misc"
)

type CallbackQuery struct {
	Prefix   string
	LevelReq AccessLevel
	ChatType ChatType
	Handler  func(*bot.Server, *gotgbot.Bot, *ext.Context) error
}

func GetCallbackQueryList() []CallbackQuery {
	CallbackQueries := make([]CallbackQuery, 0)

	// Add your callback queries here
	CallbackQueries = append(CallbackQueries, CallbackQuery{
		Prefix:   "setLang:",
		LevelReq: AccessLevelUser,
		ChatType: ChatTypePrivate,
		Handler:  misc.HandleSetLanguageCallback,
	})

	return CallbackQueries
}
