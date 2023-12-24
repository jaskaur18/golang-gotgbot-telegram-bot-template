package conv

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/conversation"
	"github.com/gotgbot/keyboard"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers/commands"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers/misc"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/menus"
	"github.com/jaskaur18/moimoiStoreBot/helpers"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

var (
	EnterPhone  = "enterPhone"
	EnterGender = "enterGender"
)

func handlePhone(b *gotgbot.Bot, c *ext.Context) error {
	if c.EffectiveMessage.Contact == nil {
		phKey := new(keyboard.Keyboard).RequestContact("Send your phone number")

		_, err := b.SendMessage(c.EffectiveChat.Id, "Incorrect!!! Share Your Phone Number Via \"Send Phone Number Button\"", &gotgbot.SendMessageOpts{
			ReplyMarkup: phKey.Build(),
		})

		if err != nil {
			_ = misc.ErrorHandler(b, c, err)
		}
	}

	ct := c.EffectiveMessage.Contact

	session, err := helper.GetSession(c.EffectiveUser.Id)
	if err != nil {
		_ = misc.ErrorHandler(b, c, err)
		return handlers.EndConversation()
	}

	session.PhoneNumber = ct.PhoneNumber

	err = session.Save()
	if err != nil {
		_ = misc.ErrorHandler(b, c, err)
		return handlers.EndConversation()
	}

	_, _ = c.EffectiveMessage.Reply(b, "Finally! Select Your Gender", &gotgbot.SendMessageOpts{
		ReplyMarkup: menus.GenderKeyboard,
	})

	return handlers.NextConversationState(EnterGender)
}

func handleGender(b *gotgbot.Bot, c *ext.Context) error {
	gender := c.EffectiveMessage.Text

	if gender != "Male" && gender != "Female" && gender != "Prefer Not To Share" {
		_, _ = c.EffectiveMessage.Reply(b, "Incorrect Select Valid Option", &gotgbot.SendMessageOpts{
			ReplyMarkup: menus.GenderKeyboard,
		})

		return nil
	}

	session, err := helper.GetSession(c.EffectiveUser.Id)
	if err != nil {
		_ = misc.ErrorHandler(b, c, err)
		return handlers.EndConversation()
	}

	if session.PhoneNumber == "" {
		_, _ = c.EffectiveMessage.Reply(b, "Error Occurred Fetching UserState Try Again", nil)
		return handlers.EndConversation()
	}

	collection, err := helpers.PB.Dao().FindCollectionByNameOrId("tgUser")
	if err != nil {
		_ = misc.ErrorHandler(b, c, err)
		return handlers.EndConversation()
	}

	record := models.NewRecord(collection)

	form := forms.NewRecordUpsert(helpers.PB, record)

	if session.ReferralCode == 0 {
		err = form.LoadData(map[string]any{
			"name":     fmt.Sprintf("%s %s", c.EffectiveUser.FirstName, c.EffectiveUser.LastName),
			"tgId":     c.EffectiveUser.Id,
			"username": c.EffectiveUser.Username,
			"gender":   gender,
			"phone":    session.PhoneNumber,
		})
	} else {
		err = form.LoadData(map[string]any{
			"name":           fmt.Sprintf("%s %s", c.EffectiveUser.FirstName, c.EffectiveUser.LastName),
			"tgId":           c.EffectiveUser.Id,
			"username":       c.EffectiveUser.Username,
			"gender":         gender,
			"phone":          session.PhoneNumber,
			"referralParent": session.ReferralCode,
		})
	}

	if err != nil {
		_, _ = c.EffectiveMessage.Reply(b, "Error Occurred Validating Data Try Again", nil)
		return handlers.EndConversation()
	}

	if err := form.Submit(); err != nil {
		_, _ = c.EffectiveMessage.Reply(b, "Error Occurred Saving Data Try Again", nil)
		return handlers.EndConversation()
	}

	_, _ = c.EffectiveMessage.Reply(b, "You Have Successfully Joined Bot", &gotgbot.SendMessageOpts{})
	_, _ = c.EffectiveMessage.Reply(b, "Welcome To The Bot", &gotgbot.SendMessageOpts{
		ReplyMarkup: menus.HomeKeyboard,
	})

	return handlers.EndConversation()
}

var UserJoin = handlers.NewConversation(
	[]ext.Handler{handlers.NewCommand("start", commands.CommandStart)},
	map[string][]ext.Handler{
		EnterPhone:  {handlers.NewMessage(literalyAnything, handlePhone)},
		EnterGender: {handlers.NewMessage(noCommands, handleGender)},
	},
	&handlers.ConversationOpts{
		Exits:        []ext.Handler{handlers.NewMessage(CancelMsg, HandleCancel)},
		StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
		AllowReEntry: true,
	},
)
