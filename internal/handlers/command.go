package handlers

import (
	"bot/cmd/bot"
	"bot/internal/handlers/commands"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type Commands struct {
	Name     string
	LevelReq AccessLevel
	ChatType ChatType
	Handler  func(*bot.Server, *gotgbot.Bot, *ext.Context) error
}

var CommandsList = []Commands{
	{
		Name:     "start",
		LevelReq: AccessLevelUser,
		Handler:  commands.CommandStart,
		ChatType: ChatTypePrivate,
	},
	{
		Name:     "lang",
		LevelReq: AccessLevelUser,
		Handler:  commands.HandleLanguage,
		ChatType: ChatTypePrivate,
	},
	{
		Name:     "admin",
		LevelReq: AccessLevelSudoAdmin,
		Handler:  commands.HandleAdmin,
		ChatType: ChatTypePrivate,
	},
	{
		Name:     "broadcast",
		LevelReq: AccessLevelAdmin,
		Handler:  commands.CommandBroadcast,
		ChatType: ChatTypePrivate,
	},
}
