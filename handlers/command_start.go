package handlers

import (
	"bot/handlers/commands"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

var commandStartHandler = handlers.NewCommand("start", commands.CommandStart)
