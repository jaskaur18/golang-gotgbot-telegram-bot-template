package conv

import (
	"context"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/conversation"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers/callbackQuery"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
	"github.com/jaskaur18/moimoiStoreBot/helpers"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"log"
)

var ProductStepper = "productStepper"

func handleProductStepper(b *gotgbot.Bot, c *ext.Context) error {
	session, err := helper.GetSession(c.EffectiveUser.Id)
	if err != nil {
		_, _ = c.EffectiveMessage.Reply(b, "❌ Error Getting UserState", &gotgbot.SendMessageOpts{})
		return handlers.EndConversation()
	}

	if session == nil || session.ProductID == "" || session.ProductSteps == nil {
		_, _ = c.EffectiveMessage.Reply(b, "❌ Something Went Wrong. Try Again later - no steps found for this product", &gotgbot.SendMessageOpts{})
		return handlers.EndConversation()
	}

	currentStep := &session.ProductSteps[session.ProductStepper]

	switch currentStep.Type {
	case "text":
		if c.EffectiveMessage.Text == "" {
			_, _ = c.EffectiveMessage.Reply(b, "❌ You Must Send A Text. /cancel to cancel", &gotgbot.SendMessageOpts{})
			return nil
		}
		currentStep.Value = c.EffectiveMessage.Text
	case "photo":
		if c.EffectiveMessage.Photo == nil {
			_, _ = c.EffectiveMessage.Reply(b, "❌ You Must Send A Photo. /cancel to cancel", &gotgbot.SendMessageOpts{})
			return nil
		}
		currentStep.Value = c.EffectiveMessage.Photo[len(c.EffectiveMessage.Photo)-1].FileId
	case "video":
		if c.EffectiveMessage.Video == nil {
			_, _ = c.EffectiveMessage.Reply(b, "❌ You Must Send A Video. /cancel to cancel", &gotgbot.SendMessageOpts{})
			return nil
		}
		currentStep.Value = c.EffectiveMessage.Video.FileId
	case "document":
		if c.EffectiveMessage.Document == nil {
			_, _ = c.EffectiveMessage.Reply(b, "❌ You Must Send A Document. /cancel to cancel", &gotgbot.SendMessageOpts{})
			return nil
		}
		currentStep.Value = c.EffectiveMessage.Document.FileId
	default:
		_, _ = c.EffectiveMessage.Reply(b, "❌ Unsupported MessageType Provided You Need To Send "+currentStep.Type, &gotgbot.SendMessageOpts{})
		return nil
	}

	session.ProductSteps[session.ProductStepper] = *currentStep
	session.ProductStepper++

	err = session.Save()
	if err != nil {
		_, _ = c.EffectiveMessage.Reply(b, "❌ Error Saving UserState", &gotgbot.SendMessageOpts{})
		return handlers.EndConversation()
	}

	if len(session.ProductSteps) == session.ProductStepper {
		tgUser, err := helpers.PB.Dao().FindFirstRecordByFilter("tgUser", "tgId = {:id}", dbx.Params{
			"id": c.EffectiveUser.Id,
		})

		if err != nil {
			_, _ = c.EffectiveMessage.Reply(b, "❌ Error Getting User", &gotgbot.SendMessageOpts{})
			return handlers.EndConversation()
		}

		gender := tgUser.GetString("gender")

		messageTxt := fmt.Sprintf("New Deal Request\n\n"+
			"Product: %s\n"+
			"UserId: %d\n"+
			"Gender: %s\n\n"+
			"User Answers: Question : Answer\n", session.ProductID, c.EffectiveUser.Id, gender)

		for _, step := range session.ProductSteps {
			messageTxt += fmt.Sprintf("%s : %s\n", step.Message, step.Value)
		}

		collection, err := helpers.PB.Dao().FindCollectionByNameOrId("deal")
		if err != nil {
			return err
		}

		record := models.NewRecord(collection)

		form := forms.NewRecordUpsert(helpers.PB, record)

		err = form.LoadData(map[string]any{
			"tgUser":  tgUser.Id,
			"status":  "Pending",
			"product": session.ProductID,
			"message": messageTxt,
		})

		if err != nil {
			_, err = c.EffectiveMessage.Reply(b, "❌ Error Creating Deal Failed To Validate", &gotgbot.SendMessageOpts{})
			return err
		}

		var files []*filesystem.File

		for _, step := range session.ProductSteps {
			if step.Type == "text" {
				continue
			}

			attachment := helper.GetDownloadLink(b, step.Value)
			f1, err := filesystem.NewFileFromUrl(context.TODO(), attachment)
			if err != nil {
				log.Printf("Error Creating Deal: %v", err)
				continue
			}

			files = append(files, f1)
		}

		if len(files) > 0 {
			err = form.AddFiles("attachment", files...)
			if err != nil {
				log.Printf("Error Creating Deal: %v", err)
				_, err = c.EffectiveMessage.Reply(b, "❌ Error Creating Deal Failed To Validate", &gotgbot.SendMessageOpts{})
				return err
			}
		}

		err = form.Submit()
		if err != nil {
			log.Printf("Error Creating Deal: %v", err)
			_, err = c.EffectiveMessage.Reply(b, "❌ Error Creating Deal Failed To Create", &gotgbot.SendMessageOpts{})
			return err
		}

		//	Update User Deal Open
		record, err = helpers.PB.Dao().FindRecordById("tgUser", tgUser.Id)
		if err != nil {
			_, err = c.EffectiveMessage.Reply(b, "❌ Error Creating Deal Failed To Update User", &gotgbot.SendMessageOpts{})
			return handlers.EndConversation()
		}

		form = forms.NewRecordUpsert(helpers.PB, record)

		err = form.LoadData(map[string]any{
			"dealOpen": true,
		})

		if err != nil {
			_, err = c.EffectiveMessage.Reply(b, "❌ Error Creating Deal Failed To Update User", &gotgbot.SendMessageOpts{})
			return handlers.EndConversation()
		}

		err = form.Submit()
		if err != nil {
			_, err = c.EffectiveMessage.Reply(b, "❌ Error Creating Deal Failed To Update User", &gotgbot.SendMessageOpts{})
			return handlers.EndConversation()
		}

		session.DealOpen = true

		err = session.Save()
		if err != nil {
			_, _ = c.EffectiveMessage.Reply(b, "❌ Error Saving UserState", &gotgbot.SendMessageOpts{})
		}

		_, _ = c.EffectiveMessage.Reply(b, "✅ Thank you for your submission. We will review it and get back to you soon.", &gotgbot.SendMessageOpts{})
		return handlers.EndConversation()

	} else {
		currentStep = &session.ProductSteps[session.ProductStepper]
		_, _ = c.EffectiveMessage.Reply(b, currentStep.Message, &gotgbot.SendMessageOpts{
			ParseMode: "HTML",
		})
	}

	return nil
}

var ProductStepperConv = handlers.NewConversation(
	[]ext.Handler{handlers.NewCallback(callbackquery.Prefix("buy:"), callbackQuery.HandleBuyProduct)},
	map[string][]ext.Handler{
		ProductStepper: {handlers.NewMessage(literalyAnything, handleProductStepper)},
	},
	&handlers.ConversationOpts{
		Exits:        []ext.Handler{handlers.NewMessage(CancelMsg, HandleCancel)},
		StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
		AllowReEntry: true,
	},
)
