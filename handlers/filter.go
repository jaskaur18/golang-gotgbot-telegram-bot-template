package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type TextFilter struct {
	Text    string
	Handler func(b *gotgbot.Bot, ctx *ext.Context) error
}

func FilterText(msg *gotgbot.Message) bool {
	for _, filter := range TextFilters {
		if msg.Text == filter.Text {
			return true
		}
	}
	return false
}

var TextFilters = []TextFilter{
	{},
}

func HandleTextFilters(b *gotgbot.Bot, ctx *ext.Context) error {
	for _, filter := range TextFilters {
		if ctx.EffectiveMessage.Text == filter.Text {
			return filter.Handler(b, ctx)
		}
	}
	return nil
}
