package commands

import (
	"bot/cmd/bot"
	"bot/internal/handlers/misc"
	"bot/internal/models"
	"bot/internal/utils"
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/keyboard"
	"github.com/lus/fluent.go/fluent"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CommandStart(s *bot.Server, b *gotgbot.Bot, ctx *ext.Context) error {
	c, err := models.Users(models.UserWhere.Telegramid.EQ(null.Int64From(ctx.EffectiveUser.Id))).Count(context.Background(), s.DB)
	if err != nil {
		return misc.ErrorHandler(b, ctx, err)
	}

	if c == 0 {
		u := &models.User{
			Telegramid: null.Int64From(ctx.EffectiveUser.Id),
			Firstname:  ctx.EffectiveUser.FirstName,
			Lastname:   ctx.EffectiveUser.LastName,
			Username:   null.StringFrom(ctx.EffectiveUser.Username),
		}

		err := u.Insert(context.Background(), s.DB, boil.Infer())
		if err != nil {
			return misc.ErrorHandler(b, ctx, err)
		}
	}

	msg := utils.GetMessage(s.Redis, ctx, "welcome", fluent.WithVariable("name", ctx.EffectiveUser.FirstName))

	_, err = ctx.EffectiveMessage.Reply(b, msg, &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: new(
			keyboard.Keyboard,
		).Text(
			"text",
		).RequestContact(
			"text",
		).Build(),
	})

	if err != nil {
		return misc.ErrorHandler(b, ctx, err)
	}

	return nil
}
