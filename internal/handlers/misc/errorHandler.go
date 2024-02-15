package misc

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/helper"
	"github.com/rs/zerolog/log"
	"runtime"
)

func ErrorHandler(b *gotgbot.Bot, ctx *ext.Context, err error) error {
	// Recovers From Panics
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered error in f", r)
		}
	}()

	// Get the filename and line number of the error occurrence
	_, file, line, _ := runtime.Caller(1)

	//log.Printf("Error Happened At %s:%d To User: %s (%d) \nError: %s", file, line, ctx.EffectiveUser.Username, ctx.EffectiveUser.Id, err)

	log.Debug().Stack().Err(err).
		Str("Username", ctx.EffectiveUser.Username).
		Int64("UserID", ctx.EffectiveUser.Id).
		Str("File", file).
		Int("Line", line).
		Msg("an error occurred while handling update")

	_, err = ctx.Message.Reply(b, "Error Happened", nil)
	for sID := range helper.SudoAdmins {
		msg := fmt.Sprintf("Error Happened At %s:%d To User: %s (%d) \nError: %v", file, line, ctx.EffectiveUser.Username, ctx.EffectiveUser.Id, err)
		_, _ = b.SendMessage(sID, msg, nil)
	}
	return err
}
