package callbackQuery

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/menus"
	"strings"
)

func HandleBack(b *gotgbot.Bot, c *ext.Context) error {
	d := strings.Split(c.CallbackQuery.Data, ":")
	if len(d) != 2 {
		_, err := b.AnswerCallbackQuery(c.CallbackQuery.Id, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "No Id For Back Action Found In The Callback Query",
			ShowAlert: true,
		})
		return err
	}

	switch d[1] {
	case "products":
		return productBack(b, c)
	}

	_, err := b.AnswerCallbackQuery(c.CallbackQuery.Id, &gotgbot.AnswerCallbackQueryOpts{
		Text:      "No Back Action Found In The Callback Query",
		ShowAlert: true,
	})

	return err
}

func productBack(b *gotgbot.Bot, c *ext.Context) error {
	key := menus.ServicesKeyboard()

	_, _ = c.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text: "Showing Available Services",
	})

	msg := "<code>-- Available Service List --</code>"

	//If the Content Of Both Messages Is Same Then Don't Edit The Message
	if c.EffectiveMessage.Text == msg {
		return nil
	}

	_, _, _ = c.EffectiveMessage.EditText(b, msg, &gotgbot.EditMessageTextOpts{
		ParseMode:   "html",
		ReplyMarkup: *key,
	})

	return nil
}
