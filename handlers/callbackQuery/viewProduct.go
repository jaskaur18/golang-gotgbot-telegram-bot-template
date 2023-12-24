package callbackQuery

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/keyboard"
	"github.com/jaskaur18/moimoiStoreBot/helpers"
	"github.com/pocketbase/pocketbase/apis"
	"strings"
)

func HandleViewProduct(b *gotgbot.Bot, c *ext.Context) error {
	d := strings.Split(c.CallbackQuery.Data, ":")
	if len(d) != 2 {
		_, err := b.AnswerCallbackQuery(c.CallbackQuery.Id, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "No Product ID Found Try Again",
			ShowAlert: true,
		})

		return err
	}

	pID := d[1]

	product, err := helpers.PB.Dao().FindRecordById("product", pID)
	if err != nil {
		if err = apis.NewNotFoundError(err.Error(), nil); err != nil {
			_, _ = b.AnswerCallbackQuery(c.CallbackQuery.Id, &gotgbot.AnswerCallbackQueryOpts{
				Text:      "Product Not Found",
				ShowAlert: true,
			})

			return err
		}

		_, _ = b.AnswerCallbackQuery(c.CallbackQuery.Id, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "Could Not Find Product",
			ShowAlert: true,
		})
	}

	hidden := product.GetBool("hidden")
	if hidden {
		_, _ = b.AnswerCallbackQuery(c.CallbackQuery.Id, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "Product Not Found",
			ShowAlert: true,
		})

		return nil
	}

	menu := new(keyboard.InlineKeyboard).Text("Continue", fmt.Sprintf("buy:%s", pID)).
		Row().Text("Back", "back:products").Build()

	msg := fmt.Sprintf("Product Name: <b>%s</b>\n\n%s", product.GetString("name"), product.GetString("description"))

	_, _ = b.AnswerCallbackQuery(c.CallbackQuery.Id, &gotgbot.AnswerCallbackQueryOpts{
		Text: "Viewing Product",
	})

	_, _, err = c.EffectiveMessage.EditText(b, msg, &gotgbot.EditMessageTextOpts{
		ParseMode:   "html",
		ReplyMarkup: *menu,
	})

	return nil
}
