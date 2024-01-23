package handlers

import (
	"bot/cmd/bot"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
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
