package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
)

type TextFilter struct {
	Text     string
	LevelReq AccessLevel
	ChatType ChatType
	Handler  func(bot *bot.Server, b *gotgbot.Bot, ctx *ext.Context) error
}

var TextFilters = []TextFilter{
	{},
}
