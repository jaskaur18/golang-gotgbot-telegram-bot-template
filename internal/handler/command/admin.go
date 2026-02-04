package command

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/service"
)

type AdminCommand struct {
	userSvc *service.UserService
}

func NewAdminCommand(userSvc *service.UserService) *AdminCommand {
	return &AdminCommand{userSvc: userSvc}
}

func (h *AdminCommand) Handle(b *gotgbot.Bot, ctx *ext.Context) error {
	// Middleware guarantees user is Admin here.

	args := ctx.Args()
	if len(args) < 3 {
		ctx.Message.Reply(b, "Usage: /admin <id> <true/false>", nil)
		return nil
	}

	targetID, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		ctx.Message.Reply(b, "Invalid ID", nil)
		return nil
	}

	isAdmin := strings.ToLower(args[2]) == "true"

	// Logic
	targetUser, err := h.userSvc.GetUser(context.Background(), targetID)
	if err != nil {
		ctx.Message.Reply(b, "User not found", nil)
		return nil
	}
	if targetUser == nil {
		ctx.Message.Reply(b, "User not found", nil)
		return nil
	}

	err = h.userSvc.UpdateUserType(context.Background(), targetID, isAdmin)
	if err != nil {
		ctx.Message.Reply(b, "Failed to update user", nil)
		return err
	}

	status := "normal user"
	if isAdmin {
		status = "admin"
	}
	ctx.Message.Reply(b, fmt.Sprintf("User %d is now %s", targetID, status), nil)
	return nil
}
