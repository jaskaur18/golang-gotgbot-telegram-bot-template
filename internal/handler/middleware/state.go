package middleware

import (
	"context"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/service"
)

func StateMiddleware(stateSvc *service.StateService, requiredState string) Middleware {
	return func(next func(b *gotgbot.Bot, ctx *ext.Context) error) func(b *gotgbot.Bot, ctx *ext.Context) error {
		return func(b *gotgbot.Bot, ctx *ext.Context) error {
			userState, err := stateSvc.GetState(context.Background(), ctx.EffectiveUser.Id)
			if err != nil {
				// DB error, maybe skip?
				return nil
			}

			// If requiredState is "*", allow any state (or non-empty?)
			// If requiredState is empty, it means "no state" (default).

			if userState != requiredState {
				// User is not in the required state. Skip this handler.
				// This allows falling through to other handlers if using a chain,
				// BUT gotgbot dispatcher stops at first match usually.
				// However, if we return nil here, gotgbot dispatcher might NOT continue to next handler
				// unless this is a filter. Middleware executes AFTER match.

				// CRITICAL: Middleware executes AFTER the dispatcher matched the command.
				// If we stop here, the update is "consumed" but no action taken.
				// For FSM, usually we want to match ONLY if state matches.
				// This implies FSM check should be a FILTER, not just middleware.
				// BUT, we can use middleware to enforce it and just do nothing if mismatch.
				// Meaning: "You triggered /cancel but you are not in a flow? Ignore or say 'nothing to cancel'".

				return nil
			}

			return next(b, ctx)
		}
	}
}
