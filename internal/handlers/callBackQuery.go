package handlers

import (
	"bot/cmd/bot"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type CallbackQuery struct {
	Prefix   string
	LevelReq AccessLevel
	ChatType ChatType
	Handler  func(*bot.Server, *gotgbot.Bot, *ext.Context) error
}

var CallbackQueries = []CallbackQuery{}
