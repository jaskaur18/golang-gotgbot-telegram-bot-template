package handlers

import (
	"bot/handlers/commands"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type Commands struct {
	Name    string
	Allowed Allowed
	Handler func(b *gotgbot.Bot, ctx *ext.Context) error
}

var CommandsList = []Commands{
	{
		Name:    "start",
		Allowed: User,
		Handler: commands.CommandStart,
	},
	{
		Name:    "adin",
		Allowed: User,
		Handler: commands.HandleAdmin,
	},
	{
		Name:    "broadcast",
		Allowed: Admin,
		Handler: commands.CommandBroadcast,
	},
}
