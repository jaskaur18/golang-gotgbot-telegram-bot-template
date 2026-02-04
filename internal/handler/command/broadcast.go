package command

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/service"
	"github.com/rs/zerolog/log"
)

type BroadcastCommand struct {
	userSvc *service.UserService
}

func NewBroadcastCommand(userSvc *service.UserService) *BroadcastCommand {
	return &BroadcastCommand{userSvc: userSvc}
}

func (h *BroadcastCommand) Handle(b *gotgbot.Bot, ctx *ext.Context) error {
	// Middleware checks for Admin logic.

	users, err := h.userSvc.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	msgText := strings.Join(ctx.Args()[1:], " ")
	if msgText == "" {
		_, err := ctx.Message.Reply(b, "Please enter a message", nil)
		return err
	}

	_, _ = ctx.Message.Reply(b, "Broadcasting...", nil)

	totalSend := 0
	for _, user := range users {
		tid := user.TelegramID
		_, err = b.SendMessage(tid, msgText, nil)
		if err != nil {
			log.Error().Err(err).Int64("tid", tid).Msg("Error sending broadcast message")
			continue
		}
		totalSend++
		time.Sleep(100 * time.Millisecond)
	}

	log.Info().Int("totalSend", totalSend).Msg("Broadcast to users")
	_, _ = ctx.Message.Reply(b, fmt.Sprintf("Broadcasted to %d users", totalSend), nil)
	return nil
}
