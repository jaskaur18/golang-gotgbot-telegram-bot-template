package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers/callbackQuery"
)

type CallbackQuery struct {
	Prefix  string
	Handler func(b *gotgbot.Bot, ctx *ext.Context) error
}

var CallbackQueries = []CallbackQuery{
	{
		Prefix:  "prod:",
		Handler: callbackQuery.HandleViewProduct,
	},
	{
		Prefix:  "back:",
		Handler: callbackQuery.HandleBack,
	},
}
