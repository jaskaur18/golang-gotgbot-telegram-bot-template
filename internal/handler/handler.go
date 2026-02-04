package handler

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/domain"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handler/command"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handler/middleware"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handler/router"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/service"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	Dispatcher *ext.Dispatcher
	UserSvc    *service.UserService
	LocSvc     *service.LocalizationService
	StateSvc   *service.StateService
	Redis      *redis.Client
}

func NewHandler(d *ext.Dispatcher, userSvc *service.UserService, locSvc *service.LocalizationService, stateSvc *service.StateService, r *redis.Client) *Handler {
	return &Handler{
		Dispatcher: d,
		UserSvc:    userSvc,
		LocSvc:     locSvc,
		StateSvc:   stateSvc,
		Redis:      r,
	}
}

func (h *Handler) Builder() *router.RouteBuilder {
	return router.New(h.Redis, h.StateSvc).
		Use(middleware.UserMiddleware(h.UserSvc))
}

func (h *Handler) Register() {
	// Commands
	startCmd := command.NewStartCommand(h.UserSvc, h.LocSvc)
	h.Dispatcher.AddHandler(handlers.NewCommand("start",
		h.Builder().Handle(startCmd.Handle),
	))

	broadcastCmd := command.NewBroadcastCommand(h.UserSvc)
	h.Dispatcher.AddHandler(handlers.NewCommand("broadcast",
		h.Builder().
			WithRole(domain.UserTypeAdmin).
			Handle(broadcastCmd.Handle),
	))

	adminCmd := command.NewAdminCommand(h.UserSvc)
	h.Dispatcher.AddHandler(handlers.NewCommand("admin",
		h.Builder().
			WithRole(domain.UserTypeAdmin).
			Handle(adminCmd.Handle),
	))

	langCmd := command.NewLangCommand(h.LocSvc)
	h.Dispatcher.AddHandler(handlers.NewCommand("lang",
		h.Builder().
			PrivateOnly().
			RateLimit(5, time.Minute). // Add Rate Limit example
			Handle(langCmd.Handle),
	))

	// Callbacks
	langCallback := command.NewLangCallback(h.LocSvc)
	h.Dispatcher.AddHandler(handlers.NewCallback(callbackquery.Prefix("setLang:"),
		h.Builder().Handle(langCallback.Handle),
	))
}
