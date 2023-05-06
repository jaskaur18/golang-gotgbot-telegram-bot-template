package commands

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/keyboard"
)

func CommandStart(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.Message.Reply(b, "Hey!", &gotgbot.SendMessageOpts{
		ReplyMarkup: new(
			keyboard.Keyboard,
		).Text(
			"text",
		).RequestContact(
			"text",
		).Build(),
	})
	return err
}
