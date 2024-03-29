package commands

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/utils"
	"github.com/rs/zerolog/log"
)

func HandleLanguage(s *bot.Server, b *gotgbot.Bot, ctx *ext.Context) error {
	args := ctx.Args()

	langArgumentNo := 2
	if len(args) < langArgumentNo {
		_, err := ctx.Message.Reply(b, s.Locale.GetMessage(s.Redis, ctx, "invalidLangArgs"), &gotgbot.SendMessageOpts{})

		return err
	}

	lang := args[1]

	if !s.Locale.CheckLocale(lang) {
		_, err := ctx.Message.Reply(b, s.Locale.GetMessage(s.Redis, ctx, "invalidLang"), &gotgbot.SendMessageOpts{})
		return err
	}

	session, err := utils.GetSession(s.Redis, ctx.EffectiveUser.Id)
	if err != nil {
		_, err := ctx.Message.Reply(b, s.Locale.GetMessage(s.Redis, ctx, "getState"), &gotgbot.SendMessageOpts{})
		return err
	}

	session.Language = lang
	err = session.Save()

	if err != nil {
		log.Error().Err(err).Stack().Msg("Failed to save session")
		_, err := ctx.Message.Reply(b, s.Locale.GetMessage(s.Redis, ctx, "saveState"), &gotgbot.SendMessageOpts{})
		return err
	}

	log.Debug().
		Int64("user", ctx.EffectiveUser.Id).
		Str("lang", lang).Msg("Language changed")

	_, err = ctx.Message.Reply(b, s.Locale.GetMessage(s.Redis, ctx, "langSuccess"), &gotgbot.SendMessageOpts{})

	return err
}
