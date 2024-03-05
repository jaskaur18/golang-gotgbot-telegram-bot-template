package misc

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/constant"
	"github.com/rs/zerolog/log"
)

func ErrorHandler(b *gotgbot.Bot, ctx *ext.Context, err error) error {
	// Recovers From Panics
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered error in f %v", r)
		}
	}()

	// Get the filename and line number of the error occurrence
	_, file, line, _ := runtime.Caller(1)

	log.Debug().Stack().Err(err).
		Str("Username", ctx.EffectiveUser.Username).
		Int64("UserID", ctx.EffectiveUser.Id).
		Str("File", file).
		Int("Line", line).
		Msg("an error occurred while handling update")

	_, err = ctx.Message.Reply(b, "Error Happened", nil)
	for sID := range constant.GetSudoAdmins() {
		msg := fmt.Sprintf("Error Happened At %s:%d To User: %s (%d) \nError: %v",
			file, line, ctx.EffectiveUser.Username, ctx.EffectiveUser.Id, err)
		_, _ = b.SendMessage(sID, msg, nil)
	}
	return err
}

func SessionError(b *gotgbot.Bot, ctx *ext.Context, err error) error {
	msg := "❌ Error Getting UserState"

	if strings.Contains(err.Error(), "error saving session to redis telegramId") {
		msg = "❌ Error Saving UserState"
	}

	log.Debug().
		Int64("TGID", ctx.EffectiveUser.Id).
		Str("Username", ctx.EffectiveUser.Username).
		Err(err).
		Msg("Error Getting Session")

	_, _ = ctx.EffectiveMessage.Reply(b, msg, nil)
	return nil
}
