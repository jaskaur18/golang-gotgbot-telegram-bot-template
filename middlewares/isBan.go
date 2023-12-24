package middlewares

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
)

func HandleBan(b *gotgbot.Bot, c *ext.Context) error {
	_, err := c.EffectiveMessage.Reply(b, "You are banned from using this bot", &gotgbot.SendMessageOpts{})
	return err
}

func IsBan(msg *gotgbot.Message) bool {
	_, ok := helper.BanUsers[msg.From.Id]
	return ok
}
