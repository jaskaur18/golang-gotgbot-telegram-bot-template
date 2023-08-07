package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type CallbackQuery struct {
	Prefix  string
	Handler func(b *gotgbot.Bot, ctx *ext.Context) error
}

var CallbackQueries = []CallbackQuery{}
