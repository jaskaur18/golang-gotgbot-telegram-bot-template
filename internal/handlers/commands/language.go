package commands

import (
	"bot/cmd/bot"
	"bot/internal/utils"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/rs/zerolog/log"
)

func HandleLanguage(s *bot.Server, b *gotgbot.Bot, ctx *ext.Context) error {
	args := ctx.Args()

	if len(args) < 2 {
		_, err := ctx.Message.Reply(b, utils.GetMessage(s.Redis, ctx, "invalidLangArgs"), &gotgbot.SendMessageOpts{})

		return err
	}

	lang := args[1]

	if !utils.CheckLocale(lang) {
		_, err := ctx.Message.Reply(b, utils.GetMessage(s.Redis, ctx, "invalidLang"), &gotgbot.SendMessageOpts{})
		return err
	}

	session, err := utils.GetSession(s.Redis, ctx.EffectiveUser.Id)
	if err != nil {
		_, err := ctx.Message.Reply(b, utils.GetMessage(s.Redis, ctx, "getState"), &gotgbot.SendMessageOpts{})
		return err
	}

	session.Language = lang
	err = session.Save()

	if err != nil {
		log.Error().Err(err).Stack().Msg("Failed to save session")
		_, err := ctx.Message.Reply(b, utils.GetMessage(s.Redis, ctx, "saveState"), &gotgbot.SendMessageOpts{})
		return err
	}

	log.Debug().
		Int64("user", ctx.EffectiveUser.Id).
		Str("lang", lang).Msg("Language changed")

	_, err = ctx.Message.Reply(b, utils.GetMessage(s.Redis, ctx, "langSuccess"), &gotgbot.SendMessageOpts{})

	return err
}
