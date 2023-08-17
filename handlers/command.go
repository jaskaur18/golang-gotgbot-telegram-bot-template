package handlers

import (
	"bot/handlers/commands"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type Commands struct {
	Name     string
	LevelReq AccessLevel
	Handler  func(b *gotgbot.Bot, ctx *ext.Context) error
}

var CommandsList = []Commands{
	{
		Name:     "start",
		LevelReq: User,
		Handler:  commands.CommandStart,
	},
	{
		Name:     "adin",
		LevelReq: User,
		Handler:  commands.HandleAdmin,
	},
	{
		Name:     "broadcast",
		LevelReq: Admin,
		Handler:  commands.CommandBroadcast,
	},
}
