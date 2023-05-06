package commands

import (
	"bot/model"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/keyboard"
)

func CommandStart(b *gotgbot.Bot, ctx *ext.Context) error {
	user := model.User{
		TelegramId: ctx.EffectiveUser.Id,
		FirstName:  ctx.EffectiveUser.FirstName,
		LastName:   ctx.EffectiveUser.LastName,
		Username:   ctx.EffectiveUser.Username,
	}
	err := model.CreateUser(&user)
	if err != nil {
		ctx.Message.Reply(b, "Error Happened", nil)
		return err
	}

	_, err = ctx.Message.Reply(b, "Hey!", &gotgbot.SendMessageOpts{
		ReplyMarkup: new(
			keyboard.Keyboard,
		).Text(
			"text",
		).RequestContact(
			"text",
		).Build(),
	})
	return nil
}
