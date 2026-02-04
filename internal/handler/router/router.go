package router

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/domain"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handler/middleware"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/service"
	"github.com/redis/go-redis/v9"
)

type RouteBuilder struct {
	middlewares []middleware.Middleware
	redis       *redis.Client
	stateSvc    *service.StateService
}

// Need to inject dependencies to Builder to use them in shortcuts
func New(r *redis.Client, s *service.StateService) *RouteBuilder {
	return &RouteBuilder{
		middlewares: []middleware.Middleware{},
		redis:       r,
		stateSvc:    s,
	}
}

// Use adds a raw middleware
func (rb *RouteBuilder) Use(m middleware.Middleware) *RouteBuilder {
	rb.middlewares = append(rb.middlewares, m)
	return rb
}

// WithRole requires a specific user role
func (rb *RouteBuilder) WithRole(role domain.UserType) *RouteBuilder {
	rb.middlewares = append(rb.middlewares, middleware.RoleMiddleware(role))
	return rb
}

// PrivateOnly requires private chat
func (rb *RouteBuilder) PrivateOnly() *RouteBuilder {
	rb.middlewares = append(rb.middlewares, middleware.ChatTypeMiddleware(true))
	return rb
}

// RateLimit adds rate limiting (e.g., 5 requests per 1 minute)
func (rb *RouteBuilder) RateLimit(limit int, window time.Duration) *RouteBuilder {
	if rb.redis != nil {
		rb.middlewares = append(rb.middlewares, middleware.RateLimitMiddleware(rb.redis, limit, window))
	}
	return rb
}

// WithState requires the user to be in a specific conversation state
func (rb *RouteBuilder) WithState(state string) *RouteBuilder {
	if rb.stateSvc != nil {
		rb.middlewares = append(rb.middlewares, middleware.StateMiddleware(rb.stateSvc, state))
	}
	return rb
}

// Handle finalizes the builder and returns the handler
func (rb *RouteBuilder) Handle(handler func(b *gotgbot.Bot, ctx *ext.Context) error) func(b *gotgbot.Bot, ctx *ext.Context) error {
	final := handler
	for i := len(rb.middlewares) - 1; i >= 0; i-- {
		final = rb.middlewares[i](final)
	}
	return final
}
