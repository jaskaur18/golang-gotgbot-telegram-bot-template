package misc

import (
	"bot/helper"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
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

	log.Printf("Error Happened At %s:%d To User: %s (%d) \nError: %s", file, line, ctx.EffectiveUser.Username, ctx.EffectiveUser.Id, err)

	_, err = ctx.Message.Reply(b, "Error Happened", nil)
	for sID := range helper.SudoAdmins {
		msg := fmt.Sprintf("Error Happened At %s:%d To User: %s (%d) \nError: %v", file, line, ctx.EffectiveUser.Username, ctx.EffectiveUser.Id, err)
		_, _ = b.SendMessage(sID, msg, nil)
	}
	return err
}
