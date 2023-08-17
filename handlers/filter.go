package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type TextFilter struct {
	Text     string
	LevelReq AccessLevel
	Handler  func(b *gotgbot.Bot, ctx *ext.Context) error
}

var TextFilters = []TextFilter{
	{},
}
