package commands

import (
	"bot/cmd/bot"
	"bot/internal/handlers/misc"
	"bot/internal/models"
	"context"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

func CommandBroadcast(s *bot.Server, b *gotgbot.Bot, ctx *ext.Context) error {
	allUsers, err := models.Users().All(context.Background(), s.DB)
	if err != nil {
		return misc.ErrorHandler(b, ctx, err)
	}

	msg := strings.Join(ctx.Args()[1:], " ")
	if msg == "" {
		_, err := ctx.Message.Reply(b, "Please enter a message", nil)
		return err
	}

	_, _ = ctx.Message.Reply(b, "Broadcasting...", nil)

	totalSend := 0
	for _, user := range allUsers {
		_, err = b.SendMessage(user.Telegramid.Int64, msg, nil)
		if err != nil {
			log.Printf("Error sending message to %d: %s", user.Telegramid.Int64, err)
			continue
		}
		totalSend++
		time.Sleep(100 * time.Millisecond)
	}

	log.Info().Int("totalSend", totalSend).Msg("Broadcast to users")

	msg = fmt.Sprintf("Broadcasted to %d users", totalSend)
	_, _ = ctx.Message.Reply(b, msg, nil)

	return nil
}
