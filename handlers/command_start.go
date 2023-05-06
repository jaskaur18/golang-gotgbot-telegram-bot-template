package handlers

import (
	"bot/handlers/commands"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type Commands struct {
	Name    string
	Handler func(b *gotgbot.Bot, ctx *ext.Context) error
}

var CommandsList = []Commands{
	{
		Name:    "start",
		Handler: commands.CommandStart,
	},
	{
		Name:    "broadcast",
		Handler: commands.CommandBroadcast,
	},
}
