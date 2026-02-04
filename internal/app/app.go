package app

import (
	"context"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/config"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handler"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/pkg/logger"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/repository/postgres"
	redisRepo "github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/repository/redis"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/service"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type App struct {
	Updater    *ext.Updater
	Dispatcher *ext.Dispatcher
	Bot        *gotgbot.Bot
	Config     *config.Config
	DB         *pgxpool.Pool
	Redis      *redis.Client
}

func Run(cfg *config.Config) {
	// Logger
	logger.Setup(cfg.Logger.Level, cfg.Logger.PrettyPrintConsole)

	// DB
	dbPool, err := pgxpool.New(context.Background(), cfg.ConnectionString())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to init DB")
	}
	defer dbPool.Close()

	// Redis
	opt, err := redis.ParseURL(cfg.Misc.RedisURI)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse Redis URI")
	}
	rClient := redis.NewClient(opt)
	defer rClient.Close()

	// Repositories
	userRepo := postgres.NewUserRepository(dbPool)
	sessionRepo := redisRepo.NewSessionRepository(rClient)

	// Services
	userSvc := service.NewUserService(userRepo)
	locSvc := service.NewLocalizationService(cfg.Misc.LocalesDir, sessionRepo)
	stateSvc := service.NewStateService(rClient)

	// Bot
	b, err := gotgbot.NewBot(cfg.Bot.Token, &gotgbot.BotOpts{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create bot")
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Error().Err(err).Msg("Handler error")
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	updater := ext.NewUpdater(dispatcher, nil)

	// Handlers
	h := handler.NewHandler(dispatcher, userSvc, locSvc, stateSvc, rClient)
	h.Register()

	// Start
	log.Info().Str("username", b.User.Username).Msg("Bot starting...")
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start polling")
	}

	updater.Idle()
}
