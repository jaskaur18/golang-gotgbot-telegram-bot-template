package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func Load(dp *ext.Dispatcher) {
	// dp.AddHandler(commandAHandler)
	// dp.AddHandler(commandBHandler)
	for _, command := range CommandsList {
		cmd := handlers.NewCommand(command.Name, command.Handler)
		dp.AddHandler(cmd)
	}

	// dp.AddHandler(inlineQueryAHandler)
	// dp.AddHandler(inlineQueryBHandler)
}
