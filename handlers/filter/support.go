package filter

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/keyboard"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
)

func HandleSupport(b *gotgbot.Bot, ctx *ext.Context) error {
	key := new(keyboard.InlineKeyboard).Url("☎ Support Bot ☎", fmt.Sprint("https://t.me/", helper.SupportBotUsername))

	_, _ = ctx.EffectiveMessage.Reply(b, "Contact our support bot for any queries regarding order or any other issue.",
		&gotgbot.SendMessageOpts{
			ReplyMarkup: key.Build(),
		})

	return nil
}
