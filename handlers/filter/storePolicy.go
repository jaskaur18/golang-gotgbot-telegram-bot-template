package filter

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
)

func HandleStorePolicy(b *gotgbot.Bot, ctx *ext.Context) error {
	msg := fmt.Sprintf(`🔏If You are dealing with us that mean's you are accepting the below T&C🗒✏️

  ●Sharing Phone Number Asking Driver Number, and Any Personal Details is Not Allowed. 

  ●If you cheat, then we'll ban you straight away

  If Any Issue Contact <a href="t.me/%s">Here</a>.`, helper.SupportGuyUsername)
	_, _ = ctx.EffectiveMessage.Reply(b, msg, &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})

	return nil
}
