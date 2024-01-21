package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers/commands"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type Commands struct {
	Name     string
	LevelReq AccessLevel
	Handler  func(b *gotgbot.Bot, ctx *ext.Context) error
}

var CommandsList = []Commands{

	{
		Name:     "admin",
		LevelReq: User,
		Handler:  commands.HandleAdmin,
	},
	{
		Name:     "broadcast",
		LevelReq: Admin,
		Handler:  commands.CommandBroadcast,
	},
}
