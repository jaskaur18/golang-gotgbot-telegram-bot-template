package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
)

type Allowed int

const (
	SudoAdmin Allowed = iota
	Admin
	User
)

func Load(dp *ext.Dispatcher) {

	for _, command := range CommandsList {
		cmd := handlers.NewCommand(command.Name, command.Handler)
		dp.AddHandler(cmd)
	}

	//CallbackQueries
	for _, callback := range CallbackQueries {
		cb := handlers.NewCallback(callbackquery.Prefix(callback.Prefix), callback.Handler)
		dp.AddHandler(cb)
	}

	//dp.AddHandler(handlers.NewMessage(MessageFilter, TextHandlers))

}

//func MessageFilter(msg *gotgbot.Message) bool {
//	return msg.Text != ""
//}
//
//func TextHandlers(b *gotgbot.Bot, c *ext.Context) error {
//	text := c.EffectiveMessage.Text
//	for _, command := range CommandsList {
//		if command.Name == text {
//			if command.Allowed == Admin && !middlewares.IsAdmin(text) {
//				return notAdminHandler(b, c)
//			} else if command.Allowed == SudoAdmin && !middlewares.IsSudoAdmin(msg) {
//				return notSudoAdminHandler
//			}
//			return command.Handler
//		}
//	}
//
//	return nil
//}
//
//func notAdminHandler(b *gotgbot.Bot, ctx *ext.Context) error {
//	_, err := ctx.EffectiveMessage.Reply(b, "You are not admin, only admin can use this command", nil)
//	return err
//}
//
//func notSudoAdminHandler(b *gotgbot.Bot, ctx *ext.Context) error {
//	_, err := ctx.EffectiveMessage.Reply(b, "You are not sudo admin, only sudo admin can use this command", nil)
//	return err
//}
