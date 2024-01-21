package handlers

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
<<<<<<< HEAD
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers/conv"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/middlewares"
	"strings"
=======
>>>>>>> parent of dc24b0d (Update i18n implementation, libraries and installation script)
)

type AccessLevel int

const (
	SudoAdmin AccessLevel = iota
	Admin
	User
)

func LoadHandlers(dp *ext.Dispatcher) {
	//Ban Check
	dp.AddHandler(handlers.NewMessage(middlewares.IsBan, middlewares.HandleBan))
	dp.AddHandler(handlers.NewMessage(middlewares.IsDealOpen, middlewares.DealHandler))
	dp.AddHandler(conv.UserJoin)
	dp.AddHandler(conv.ProductStepperConv)
	loadCallbackQueryHandlers(dp)
	loadMessageFilterHandler(dp)
}

func loadCallbackQueryHandlers(dp *ext.Dispatcher) {
	for _, callback := range CallbackQueries {
		callbackHandler := handlers.NewCallback(callbackquery.Prefix(callback.Prefix), callback.Handler)
		dp.AddHandler(callbackHandler)
	}
}

func loadMessageFilterHandler(dp *ext.Dispatcher) {
	msgFilterHandler := handlers.NewMessage(TextMessageFilter, HandleTextMessageFilter)
	dp.AddHandler(msgFilterHandler)
}

func TextMessageFilter(msg *gotgbot.Message) bool {
	return msg.Text != ""
}

func HandleTextMessageFilter(b *gotgbot.Bot, c *ext.Context) error {
	if message.Command(c.EffectiveMessage) {
		return handleCommand(b, c)
	}

	for _, filter := range TextFilters {
		if filter.Text == c.EffectiveMessage.Text {
			if filter.LevelReq == Admin && !middlewares.IsAdmin(c.EffectiveMessage) {
				return handleNotAllowed(b, c, "admin")
			} else if filter.LevelReq == SudoAdmin && !middlewares.IsSudoAdmin(c.EffectiveMessage) {
				return handleNotAllowed(b, c, "sudo admin")
			}
			return filter.Handler(b, c)
		}
	}

	return nil
}

func handleCommand(b *gotgbot.Bot, c *ext.Context) error {
	text := c.EffectiveMessage.Text
	for _, cmd := range CommandsList {
		if cmd.Name == text {
			if cmd.LevelReq == Admin && !middlewares.IsAdmin(c.EffectiveMessage) {
				return handleNotAllowed(b, c, "admin")
			} else if cmd.LevelReq == SudoAdmin && !middlewares.IsSudoAdmin(c.EffectiveMessage) {
				return handleNotAllowed(b, c, "sudo admin")
			}
			return cmd.Handler(b, c)
		}
	}
	return nil
}

func handleNotAllowed(b *gotgbot.Bot, c *ext.Context, role string) error {
	msg := fmt.Sprintf("You are not %s, only %s can use this command", role, role)
	_, err := c.EffectiveMessage.Reply(b, msg, nil)
	return err
}
