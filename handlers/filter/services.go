package filter

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers/misc"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/menus"
	"github.com/jaskaur18/moimoiStoreBot/helpers"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
)

func HandleServices(b *gotgbot.Bot, c *ext.Context) error {
	user, err := helpers.PB.Dao().FindFirstRecordByFilter("tgUser", "tgId = {:id}",
		dbx.Params{"id": c.EffectiveUser.Id})

	if err != nil {
		if apis.NewNotFoundError(err.Error(), nil) != nil {
			_, _ = c.EffectiveMessage.Reply(b, "Registration Required!!! Please Start Bot First - /start", nil)
			return nil
		}
	}

	session, err := helper.GetSession(c.EffectiveUser.Id)
	if err != nil {
		_, _ = c.EffectiveMessage.Reply(b, "Failed To Get UserState", nil)
		return nil
	}

	dealOpen := user.Get("dealOpen").(bool)
	if dealOpen {
		if !session.DealOpen {
			session.DealOpen = true
			_ = session.Save()
		}

		_, _ = c.EffectiveMessage.Reply(b, "You have an ongoing deal", nil)
		return nil
	}

	key := menus.ServicesKeyboard()

	_, err = c.Message.Reply(b, "<code>-- Available Service List --</code>", &gotgbot.SendMessageOpts{
		ParseMode:   "html",
		ReplyMarkup: key,
	})

	if err != nil {
		return misc.ErrorHandler(b, c, err)
	}

	return nil
}
