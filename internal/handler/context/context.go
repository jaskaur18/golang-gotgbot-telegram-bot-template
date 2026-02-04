package context

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/domain"
)

const (
	keyUser   = "user"
	keyLocale = "locale"
)

// SetUser stores the domain User in the context
func SetUser(ctx *ext.Context, user *domain.User) {
	if ctx.Data == nil {
		ctx.Data = make(map[string]interface{})
	}
	ctx.Data[keyUser] = user
}

// GetUser retrieves the domain User from the context
func GetUser(ctx *ext.Context) *domain.User {
	if ctx.Data == nil {
		return nil
	}
	if v, ok := ctx.Data[keyUser]; ok {
		return v.(*domain.User)
	}
	return nil
}

// SetLocale stores the resolved locale string
func SetLocale(ctx *ext.Context, locale string) {
	if ctx.Data == nil {
		ctx.Data = make(map[string]interface{})
	}
	ctx.Data[keyLocale] = locale
}

// GetLocale retrieves the locale string
func GetLocale(ctx *ext.Context) string {
	if ctx.Data == nil {
		return "en"
	}
	if v, ok := ctx.Data[keyLocale]; ok {
		return v.(string)
	}
	return "en"
}
