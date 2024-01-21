package commands

import (
<<<<<<< HEAD
	"errors"
	"fmt"
=======
	"bot/db"
	"bot/handlers/misc"
	"bot/helper"
	"context"
>>>>>>> parent of dc24b0d (Update i18n implementation, libraries and installation script)
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgbot/keyboard"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers/misc"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/menus"
	"github.com/jaskaur18/moimoiStoreBot/helpers"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"strings"
)

func CommandStart(b *gotgbot.Bot, c *ext.Context) error {
	_, err := helpers.PB.Dao().FindFirstRecordByFilter("tgUser", "tgId = {:id}", dbx.Params{
		"id": c.EffectiveUser.Id,
	})

<<<<<<< HEAD
	if err != nil {
		if apis.NewNotFoundError("", err) != nil {
			_, err := helpers.PB.Dao().FindRecordsByFilter("tgUser", "", "created", 0, 0, dbx.Params{})
			if err != nil && apis.NewNotFoundError("", err) != nil {
				session, err := helper.GetSession(c.EffectiveUser.Id)
				if err != nil {
					_ = misc.ErrorHandler(b, c, err)
				}

				session.ReferralCode = 0
				err = session.Save()
				if err != nil {
					_ = misc.ErrorHandler(b, c, err)
				}

			} else {
				if !CheckUserReferralCode(c, b) {
					return nil
				}
			}

			phKey := new(keyboard.Keyboard).
				RequestContact("Send Phone Number").Build()

			_, err = c.Message.Reply(b, "To Continue Please Share Your Phone Number Via \"Send Phone Number Button\"",
				&gotgbot.SendMessageOpts{
					ReplyMarkup: phKey,
				})

			return handlers.NextConversationState("enterPhone")
		}
	}

	_, err = c.Message.Reply(b, "<code>Welcome</code>", &gotgbot.SendMessageOpts{
		ParseMode:   "html",
		ReplyMarkup: menus.HomeKeyboard,
	})

	if err != nil {
		return misc.ErrorHandler(b, c, err)
	}

	key := menus.ServicesKeyboard()

	_, err = c.Message.Reply(b, "<code>-- Available Service List --</code>", &gotgbot.SendMessageOpts{
		ParseMode:   "html",
		ReplyMarkup: key,
=======
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
>>>>>>> parent of dc24b0d (Update i18n implementation, libraries and installation script)
	})

	if err != nil {
		return misc.ErrorHandler(b, c, err)
	}

	return nil
}

func CheckUserReferralCode(c *ext.Context, b *gotgbot.Bot) bool {
	var match string
	data := strings.Split(c.EffectiveMessage.Text, " ")

	if len(data) > 1 {
		match = data[1]
	}

	if match != "" {
		_, err := helpers.PB.Dao().FindFirstRecordByFilter("tgUser", "tgId = {:code}", dbx.Params{
			"code": match,
		})

		if err != nil {
			if apis.NewNotFoundError("", err) != nil {
				_, _ = c.EffectiveMessage.Reply(b, "Invalid Referral Link", nil)
				return false
			}
		}

		session, err := helper.GetSession(c.EffectiveUser.Id)
		if err != nil {
			_ = misc.ErrorHandler(b, c, err)
		}

		code, ok := helper.StringToInt64(match)
		if !ok {
			_ = misc.ErrorHandler(b, c, errors.New("invalid referral code"))
			return false
		}

		session.ReferralCode = code
		err = session.Save()
		if err != nil {
			_ = misc.ErrorHandler(b, c, err)
			return false
		}
	} else {
		key := new(keyboard.InlineKeyboard).Text("Contact Here", fmt.Sprint("https://t.me/", c.EffectiveUser.Username))

		_, err := c.EffectiveMessage.Reply(b, "Please Contact Here", &gotgbot.SendMessageOpts{
			ReplyMarkup: key.Build(),
		})
		if err != nil {
			_ = misc.ErrorHandler(b, c, err)
			return false
		}
	}
	return true
}
