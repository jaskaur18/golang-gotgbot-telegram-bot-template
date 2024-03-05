package commands

import (
	"context"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/keyboard"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/db"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handlers/misc"
	"github.com/lus/fluent.go/fluent"
)

func CommandStart(s *bot.Server, b *gotgbot.Bot, ctx *ext.Context) error {
	c, err := s.Queries.CheckUserExist(context.Background(), pgtype.Int8{Int64: ctx.EffectiveUser.Id, Valid: true})
	if err != nil {
		return misc.ErrorHandler(b, ctx, err)
	}

	if !c {
		u := db.CreateUserParams{
			TelegramID: pgtype.Int8{Int64: ctx.EffectiveUser.Id, Valid: true},
			FirstName:  ctx.EffectiveUser.FirstName,
			LastName:   pgtype.Text{String: ctx.EffectiveUser.LastName, Valid: true},
			Username:   pgtype.Text{String: ctx.EffectiveUser.Username, Valid: true},
			UserType:   db.UsertypeUSER,
		}

		_, err := s.Queries.CreateUser(context.Background(), u)
		if err != nil {
			return misc.ErrorHandler(b, ctx, err)
		}
	}

	msg := s.Locale.GetMessage(s.Redis, ctx, "welcome", fluent.WithVariable("name", ctx.EffectiveUser.FirstName))

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
