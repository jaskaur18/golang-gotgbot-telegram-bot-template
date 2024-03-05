package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handlers/commands"
)

type Commands struct {
	Name     string
	LevelReq AccessLevel
	ChatType ChatType
	Handler  func(*bot.Server, *gotgbot.Bot, *ext.Context) error
}

func GetCommandList() []Commands {
	CommandsList := []Commands{
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

	return CommandsList
}
