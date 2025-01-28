package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handlers/commands"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handlers/misc"
)

type Commands struct {
	Name        string
	LevelReq    AccessLevel
	ChatType    ChatType
	Description string
	Handler     func(*bot.Server, *gotgbot.Bot, *ext.Context) error
}

func GetCommandList() []Commands {
	CommandsList := []Commands{
		{
			Name:        "start",
			LevelReq:    AccessLevelUser,
			Handler:     commands.CommandStart,
			ChatType:    ChatTypePrivate,
			Description: "Start the bot",
		},
		{
			Name:        "lang",
			LevelReq:    AccessLevelUser,
			Handler:     misc.HandleLanguageInline,
			ChatType:    ChatTypePrivate,
			Description: "Change the bot language",
		},
		{
			Name:        "admin",
			LevelReq:    AccessLevelSudoAdmin,
			Handler:     commands.HandleAdmin,
			ChatType:    ChatTypePrivate,
			Description: "Admin commands",
		},
		{
			Name:        "broadcast",
			LevelReq:    AccessLevelAdmin,
			Handler:     commands.CommandBroadcast,
			ChatType:    ChatTypePrivate,
			Description: "Broadcast a message to all users",
		},
	}

	return CommandsList
}
