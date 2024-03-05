package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/config"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/constant"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/db"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/handlers"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/utils"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

const (
	TimeoutSeconds = 10
)

func App() {
	// Load config
	c := config.DefaultServiceConfig()

	// Load constants
	constant.InitConstants(&c)

	// Setup logger
	utils.SetupLogger(config.LogLevelFromString(c.Logger.Level.String()), c.Logger.PrettyPrintConsole)

	var err error
	var dbInstance *pgxpool.Pool

	if dbInstance, err = bot.InitDB(&c); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}

	q := db.New(dbInstance)

	var r *redis.Client
	ctx, cancel := context.WithTimeout(context.Background(), TimeoutSeconds*time.Second)
	log.Info().Msg("Initializing redis")
	if r, err = bot.InitRedis(ctx, c.Misc.RedisURI); err != nil {
		cancel()
		log.Fatal().Err(err).Msg("Failed to initialize redis")
	}
	cancel()

	locale := utils.NewLocaleLoader(c.Misc.LocalesDir)

	s := bot.NewServer(c, dbInstance, q, r, locale)

	log.Info().Msg("Initializing bot")
	if err := s.InitBot(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize bot")
	}

	// Attach Handlers to Dispatcher
	handlers.LoadHandlers(s)

	go func() {
		if err := s.Start(); err != nil {
			log.Fatal().Err(err).Msg("Failed to start bot")
		}

		s.Updater.Idle()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// Shutdown Bot
	s.Shutdown()
}
