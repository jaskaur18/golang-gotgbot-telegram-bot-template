package commands

import (
	"bot/helper"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func HandleLanguage(b *gotgbot.Bot, c *ext.Context) error {
	args := c.Args()

	if len(args) < 2 {
		_, err := c.Message.Reply(b, "invalidArgs", &gotgbot.SendMessageOpts{
			ParseMode: "local",
		})

		return err
	}

	lang := args[1]

	if !helper.CheckLocale(lang) {
		_, err := c.Message.Reply(b, "invalidLang", &gotgbot.SendMessageOpts{
			ParseMode: "local",
		})
		return err
	}

	s, err := helper.GetSession(c.EffectiveUser.Id)
	if err != nil {
		_, err := c.Message.Reply(b, "getState", &gotgbot.SendMessageOpts{
			ParseMode: "local",
		})
		return err
	}

	s.Language = lang
	err = s.Save()

	if err != nil {
		_, err := c.Message.Reply(b, "saveState", &gotgbot.SendMessageOpts{
			ParseMode: "local",
		})
		return err
	}

	_, err = c.Message.Reply(b, "success", &gotgbot.SendMessageOpts{
		ParseMode: "local",
	})

	return err
}
