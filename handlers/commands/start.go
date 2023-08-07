package commands

import (
	"bot/db"
	"bot/handlers/misc"
	"bot/helper"
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/keyboard"
)

func CommandStart(b *gotgbot.Bot, c *ext.Context) error {

	_, err := helper.DB.User.CreateOne(
		db.User.TelegramID.Set(int(c.EffectiveUser.Id)),
		db.User.FirstName.Set(c.EffectiveUser.FirstName),
		db.User.LastName.Set(c.EffectiveUser.LastName),
		db.User.Username.Set(c.EffectiveUser.Username),
	).Exec(context.Background())

	if err != nil {
		return misc.ErrorHandler(b, c, err)
	}

	_, err = c.Message.Reply(b, "Hey!", &gotgbot.SendMessageOpts{
		ReplyMarkup: new(
			keyboard.Keyboard,
		).Text(
			"text",
		).RequestContact(
			"text",
		).Build(),
	})

	if err != nil {
		return misc.ErrorHandler(b, c, err)
	}

	return nil
}
