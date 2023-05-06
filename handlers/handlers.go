package handlers

import "github.com/PaulSonOfLars/gotgbot/v2/ext"

func Load(dp *ext.Dispatcher) {
	// dp.AddHandler(commandAHandler)
	// dp.AddHandler(commandBHandler)
	dp.AddHandler(commandStartHandler)

	// dp.AddHandler(inlineQueryAHandler)
	// dp.AddHandler(inlineQueryBHandler)
}
