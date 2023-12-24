package middlewares

import (
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
	"github.com/jaskaur18/moimoiStoreBot/helpers"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"log"
)

func DealHandler(b *gotgbot.Bot, c *ext.Context) error {
	session, err := helper.GetSession(c.EffectiveUser.Id)
	if err != nil {
		_, err = b.SendMessage(c.EffectiveChat.Id, "Failed To Get UserState", nil)
		return err
	}

	user, err := helpers.PB.Dao().FindFirstRecordByFilter("tgUser", "tgId = {:id}",
		dbx.Params{"id": c.EffectiveUser.Id})

	if err != nil {
		_, _ = c.EffectiveMessage.Reply(b, "Registration Required!!! Please Start Bot First - /start", nil)
		return nil
	}

	deal, err := helpers.PB.Dao().FindFirstRecordByFilter("deal",
		"tgUser = {:id} && (status = \"Pending\" || status = \"Accepted\")", dbx.Params{
			"id": user.Id,
		})

	if err != nil {
		log.Printf("DealHandler: %v", err)
		//session.DealOpen = false
		//_ = session.Save()

		_, err = b.SendMessage(c.EffectiveChat.Id, "Failed To Get Deal", nil)
		return err
	}

	status := deal.GetString("status")

	if status == "Pending" {
		_, err = b.SendMessage(c.EffectiveChat.Id, "Deal is pending. You will be notified when the seller accepts your deal", nil)
		return err
	}

	if status != "Accepted" {
		session.DealOpen = false
		_ = session.Save()

		_, err = b.SendMessage(c.EffectiveChat.Id, "No active deal - resetting user state", nil)
		return err
	}

	tgUser, err := helpers.PB.Dao().FindFirstRecordByFilter("tgUser", "tgId = {:id}", dbx.Params{
		"id": c.EffectiveUser.Id,
	})

	if err != nil {
		session.DealOpen = false
		_ = session.Save()

		_, err = b.SendMessage(c.EffectiveChat.Id, "Failed To Get User", nil)
		return err
	}

	var t string

	if c.EffectiveMessage.Text != "" {
		t = c.EffectiveMessage.Text
	} else if c.EffectiveMessage.Caption != "" {
		t = c.EffectiveMessage.Caption
	}

	collection, err := helpers.PB.Dao().FindCollectionByNameOrId("message")
	if err != nil {
		return err
	}

	record := models.NewRecord(collection)

	form := forms.NewRecordUpsert(helpers.PB, record)

	message := map[string]any{
		"type":   "text",
		"text":   t,
		"deal":   deal.Id,
		"seller": deal.GetString("seller"),
		"tgUser": tgUser.Id,
		"from":   "user",
	}

	attachment := ""

	if c.EffectiveMessage.Text != "" {
		message["text"] = c.EffectiveMessage.Text
	} else if c.EffectiveMessage.Photo != nil {
		message["type"] = "photo"
		attachment = helper.GetDownloadLink(b, c.EffectiveMessage.Photo[len(c.EffectiveMessage.Photo)-1].FileId)
	} else if c.EffectiveMessage.Video != nil {
		message["type"] = "video"
		attachment = helper.GetDownloadLink(b, c.EffectiveMessage.Video.FileId)
	} else if c.EffectiveMessage.Document != nil {
		message["type"] = "document"
		attachment = helper.GetDownloadLink(b, c.EffectiveMessage.Document.FileId)
	} else if c.EffectiveMessage.Animation != nil {
		attachment = helper.GetDownloadLink(b, c.EffectiveMessage.Animation.FileId)
	} else {
		_, err = b.SendMessage(c.EffectiveChat.Id, "Invalid Message Type", nil)
		return err
	}

	err = form.LoadData(message)
	if err != nil {
		log.Printf("LoadData: %v", err)
		_, err = b.SendMessage(c.EffectiveChat.Id, "Failed To Validate Message", nil)
		return err
	}

	if attachment != "" {

		f1, err := filesystem.NewFileFromUrl(context.TODO(), attachment)
		if err != nil {
			_, err = b.SendMessage(c.EffectiveChat.Id, "Failed To Create File", nil)
			return err
		}

		err = form.AddFiles("attachment", f1)
		if err != nil {
			_, err = b.SendMessage(c.EffectiveChat.Id, "Failed To Add File", nil)
			return err
		}
	}

	if err := form.Submit(); err != nil {
		_, err = b.SendMessage(c.EffectiveChat.Id, "Failed To Save Message", nil)
		return err
	}

	return nil
}

func IsDealOpen(msg *gotgbot.Message) bool {
	session, err := helper.GetSession(msg.From.Id)
	if err != nil {
		return false
	}

	if !session.DealOpen {
		return false
	}

	tgUser, err := helpers.PB.Dao().FindFirstRecordByFilter("tgUser", "tgId = {:id}", dbx.Params{
		"id": msg.From.Id,
	})

	if err != nil {
		return false
	}

	dealOpen := tgUser.GetBool("dealOpen")
	if !dealOpen {
		session.DealOpen = false
		_ = session.Save()
		return false
	}

	return true
}
