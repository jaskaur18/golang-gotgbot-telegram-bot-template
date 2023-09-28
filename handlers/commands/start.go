package commands

import (
	"bot/db"
	"bot/handlers/misc"
	"bot/helper"
	"context"
	"errors"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/keyboard"
)

func CommandStart(b *gotgbot.Bot, c *ext.Context) error {

	userExits, err := helper.DB.User.FindFirst(
		db.User.TelegramID.Equals(db.BigInt(c.EffectiveUser.Id)),
	).Exec(context.Background())

	if err != nil && !errors.Is(err, db.ErrNotFound) {
		return misc.ErrorHandler(b, c, err)
	}

	if userExits == nil {
		_, err = helper.DB.User.CreateOne(
			db.User.TelegramID.Set(db.BigInt(c.EffectiveUser.Id)),
			db.User.FirstName.Set(c.EffectiveUser.FirstName),
			db.User.LastName.Set(c.EffectiveUser.LastName),
			db.User.Username.Set(c.EffectiveUser.Username),
		).Exec(context.Background())

		if err != nil {
			return misc.ErrorHandler(b, c, err)
		}
	}

	_, err = c.Message.Reply(b, "welcome", &gotgbot.SendMessageOpts{
		ParseMode: "local,html",
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
