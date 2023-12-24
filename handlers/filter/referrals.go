package filter

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/moimoiStoreBot/helpers"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
)

func HandleReferrals(b *gotgbot.Bot, c *ext.Context) error {
	user, err := helpers.PB.Dao().FindFirstRecordByFilter("tgUser", "tgId = {:id}",
		dbx.Params{"id": c.EffectiveUser.Id})

	if err != nil {
		if apis.NewNotFoundError(err.Error(), nil) != nil {
			_, _ = c.EffectiveMessage.Reply(b, "Registration Required!!! Please Start Bot First - /start", nil)
			return nil
		}
	}

	dealOpen := user.Get("dealOpen").(bool)
	if dealOpen {
		_, _ = c.EffectiveMessage.Reply(b, "You have an ongoing deal", nil)
		return nil
	}

	referralLink := fmt.Sprintf("t.me/%s?start=%d", b.Username, c.EffectiveUser.Id)

	_, _ = c.EffectiveMessage.Reply(b, fmt.Sprintf("Your Referral link is <code>%s</code>", referralLink), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})

	return nil
}
