package command

import (
	"context"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/service"
)

type LangCommand struct {
	locSvc *service.LocalizationService
}

func NewLangCommand(locSvc *service.LocalizationService) *LangCommand {
	return &LangCommand{locSvc: locSvc}
}

func (h *LangCommand) Handle(b *gotgbot.Bot, ctx *ext.Context) error {
	locales := h.locSvc.GetAvailableLocales()
	current := h.locSvc.GetUserLocale(context.Background(), ctx.EffectiveUser.Id)

	var keyboard [][]gotgbot.InlineKeyboardButton
	for _, l := range locales {
		text := strings.ToUpper(l)
		if l == current {
			text = "✅ " + text
		}
		keyboard = append(keyboard, []gotgbot.InlineKeyboardButton{
			{Text: text, CallbackData: "setLang:" + l},
		})
	}

	_, err := ctx.Message.Reply(b, h.locSvc.GetMessage(context.Background(), ctx.EffectiveUser.Id, "langChoose"), &gotgbot.SendMessageOpts{
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: keyboard},
	})
	return err
}
