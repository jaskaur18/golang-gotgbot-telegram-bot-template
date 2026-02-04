package middleware

import (
	"context"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/domain"
	handlerContext "github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handler/context"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/service"
)

type Middleware func(next func(b *gotgbot.Bot, ctx *ext.Context) error) func(b *gotgbot.Bot, ctx *ext.Context) error

// UserMiddleware ensure the user exists in DB and loads it into context
func UserMiddleware(userSvc *service.UserService) Middleware {
	return func(next func(b *gotgbot.Bot, ctx *ext.Context) error) func(b *gotgbot.Bot, ctx *ext.Context) error {
		return func(b *gotgbot.Bot, ctx *ext.Context) error {
			user, err := userSvc.GetUser(context.Background(), ctx.EffectiveUser.Id)
			if err != nil {
				return err
			}

			if user == nil {
				err = userSvc.RegisterUser(context.Background(),
					ctx.EffectiveUser.Id,
					ctx.EffectiveUser.FirstName,
					&ctx.EffectiveUser.LastName,
					&ctx.EffectiveUser.Username,
				)
				if err != nil {
					return err
				}
				user, err = userSvc.GetUser(context.Background(), ctx.EffectiveUser.Id)
				if err != nil {
					return err
				}
			}

			handlerContext.SetUser(ctx, user)
			return next(b, ctx)
		}
	}
}

// RoleMiddleware checks if user has required role
func RoleMiddleware(requiredRole domain.UserType) Middleware {
	return func(next func(b *gotgbot.Bot, ctx *ext.Context) error) func(b *gotgbot.Bot, ctx *ext.Context) error {
		return func(b *gotgbot.Bot, ctx *ext.Context) error {
			user := handlerContext.GetUser(ctx)
			if user == nil {
				return next(b, ctx)
			}

			if requiredRole == domain.UserTypeAdmin && user.UserType != domain.UserTypeAdmin {
				ctx.Message.Reply(b, "⛔ Access Denied: Admins only.", nil)
				return nil
			}

			return next(b, ctx)
		}
	}
}

// ChatTypeMiddleware checks chat type
func ChatTypeMiddleware(isPrivate bool) Middleware {
	return func(next func(b *gotgbot.Bot, ctx *ext.Context) error) func(b *gotgbot.Bot, ctx *ext.Context) error {
		return func(b *gotgbot.Bot, ctx *ext.Context) error {
			isPrivateChat := ctx.EffectiveChat.Type == "private"
			if isPrivate && !isPrivateChat {
				ctx.Message.Reply(b, "⚠️ This command works only in private chats.", nil)
				return nil
			}
			return next(b, ctx)
		}
	}
}
