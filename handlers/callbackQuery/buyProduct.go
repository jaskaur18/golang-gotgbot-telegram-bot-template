package callbackQuery

import (
	"encoding/json"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers/misc"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
	"github.com/jaskaur18/moimoiStoreBot/helpers"
	tp "github.com/jaskaur18/moimoiStoreBot/types"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/tools/types"
	"strings"
)

func HandleBuyProduct(b *gotgbot.Bot, c *ext.Context) error {
	d := strings.Split(c.CallbackQuery.Data, ":")

	if len(d) != 2 {
		_, err := c.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "❌ Invalid Data ID Not Found In The Callback Query",
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

	prodSteps := product.Get("steps").(types.JsonRaw)

	if prodSteps == nil {
		_, _ = b.AnswerCallbackQuery(c.CallbackQuery.Id, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "Product is corrupted please report this to the admin",
			ShowAlert: true,
		})

		return nil
	}

	session, err := helper.GetSession(c.EffectiveUser.Id)
	if err != nil {
		_, err := c.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "❌ Error Getting UserState",
			ShowAlert: true,
		})
		return err
	}

	session.ProductID = pID
	session.ProductStepper = 0

	//Steps Will Be JSON Array
	//[
	//  {
	//    "id": "name",
	//    "message": "Enter Your Name Step",
	//    "type": "text",
	//    "value": ""
	//  },
	//  {
	//    "id": "age",
	//    "message": "Enter Your Age Step",
	//    "type": "text",
	//    "value": ""
	//  },
	//  {
	//    "id": "photo",
	//    "message": "Enter Your Photo Step",
	//    "type": "photo",
	//    "value": ""
	//  }
	//]

	var steps []tp.ProductSteps

	err = json.Unmarshal(prodSteps, &steps)
	if err != nil {
		return misc.ErrorHandler(b, c, err)
	}

	session.ProductSteps = steps

	err = session.Save()
	if err != nil {
		_, err := c.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      "❌ Error Saving UserState",
			ShowAlert: true,
		})
		return err
	}

	_, _ = c.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text: "✅ Product Selected",
	})

	_, _, _ = c.EffectiveMessage.EditReplyMarkup(b, nil)

	_, _, err = c.EffectiveMessage.EditText(b, "Please Enter Your Name", &gotgbot.EditMessageTextOpts{})

	return handlers.NextConversationState("productStepper")
}
