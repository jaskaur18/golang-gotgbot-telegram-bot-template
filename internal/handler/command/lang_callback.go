package command

import (
	"context"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/service"
)

type LangCallback struct {
	locSvc *service.LocalizationService
}

func NewLangCallback(locSvc *service.LocalizationService) *LangCallback {
	return &LangCallback{locSvc: locSvc}
}

func (h *LangCallback) Handle(b *gotgbot.Bot, ctx *ext.Context) error {
	data := ctx.CallbackQuery.Data
	if !strings.HasPrefix(data, "setLang:") {
		return nil
	}
	locale := strings.TrimPrefix(data, "setLang:")

	if !h.locSvc.CheckLocale(locale) {
		ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{Text: "Invalid locale"})
		return nil
	}

	err := h.locSvc.SetUserLocale(context.Background(), ctx.EffectiveUser.Id, locale)
	if err != nil {
		ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{Text: "Failed to save"})
		return err
	}

	ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text: h.locSvc.GetMessage(context.Background(), ctx.EffectiveUser.Id, "langSuccess"),
	})
	return nil
}
