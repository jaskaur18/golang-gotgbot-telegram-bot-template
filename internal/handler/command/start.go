package command

import (
	"context"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/gotgbot/keyboard"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/service"
	"github.com/lus/fluent.go/fluent"
)

type StartCommand struct {
	userSvc *service.UserService
	locSvc  *service.LocalizationService
}

func NewStartCommand(userSvc *service.UserService, locSvc *service.LocalizationService) *StartCommand {
	return &StartCommand{userSvc: userSvc, locSvc: locSvc}
}

func (h *StartCommand) Handle(b *gotgbot.Bot, ctx *ext.Context) error {
	// User registration is handled by UserMiddleware now.

	// We can get user from context if needed:
	// user := handlerContext.GetUser(ctx)

	msg := h.locSvc.GetMessageWithOptions(context.Background(), ctx.EffectiveUser.Id, "welcome", fluent.WithVariable("name", ctx.EffectiveUser.FirstName))

	_, err := ctx.EffectiveMessage.Reply(b, msg, &gotgbot.SendMessageOpts{
		ParseMode: "html",
		ReplyMarkup: new(
			keyboard.Keyboard,
		).Text(
			"Help",
		).RequestContact(
			"Send Contact",
		).Build(),
	})

	return err
}
