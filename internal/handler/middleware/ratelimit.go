package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/redis/go-redis/v9"
)

func RateLimitMiddleware(r *redis.Client, limit int, window time.Duration) Middleware {
	return func(next func(b *gotgbot.Bot, ctx *ext.Context) error) func(b *gotgbot.Bot, ctx *ext.Context) error {
		return func(b *gotgbot.Bot, ctx *ext.Context) error {
			// Simple sliding window or fixed window
			// Key: ratelimit:userID:commandhash? Or just global user rate limit?
			// Let's do per-user limit for this generic middleware
			userID := ctx.EffectiveUser.Id
			key := fmt.Sprintf("ratelimit:%d", userID)

			// Increment
			count, err := r.Incr(context.Background(), key).Result()
			if err != nil {
				// On error, fail open or close? Fail open is safer for usability.
				return next(b, ctx)
			}

			// Set expiration on first increment
			if count == 1 {
				r.Expire(context.Background(), key, window)
			}

			if count > int64(limit) {
				// Rate limited.
				// Optionally warn user once? For now, silent ignore or warn.
				if count == int64(limit)+1 {
					ctx.EffectiveMessage.Reply(b, "⚠️ Rate limit exceeded. Please wait.", nil)
				}
				return nil
			}

			return next(b, ctx)
		}
	}
}
