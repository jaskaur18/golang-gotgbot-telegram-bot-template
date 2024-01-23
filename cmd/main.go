package cmd

import (
	"bot/cmd/bot"
	"bot/internal/config"
	"bot/internal/handlers"
	"bot/internal/helper"
	"bot/internal/utils"
	"context"
	"database/sql"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func App() {

	// Load config
	c := config.DefaultServiceConfigFromEnv()

	// Load constants
	helper.InitConstants(&c)

	// Setup logger
	utils.SetupLogger(c.Logger.Level, c.Logger.PrettyPrintConsole)

	var err error
	var db *sql.DB

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	log.Info().Msg("Initializing database")
	if db, err = bot.InitDB(ctx, &c.Database); err != nil {
		cancel()
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}
	cancel()

	// Setup sql boiler Set DB Globally for all models
	boil.SetDB(db)

	var r *redis.Client
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	log.Info().Msg("Initializing redis")
	if r, err = bot.InitRedis(ctx, c.ManagementServer.RedisUri); err != nil {
		cancel()
		log.Fatal().Err(err).Msg("Failed to initialize redis")
	}
	cancel()

	s := bot.NewServer(c, db, r)

	log.Info().Msg("Initializing bot")
	if err := s.InitBot(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize bot")
	}

	//Attach Handlers to Dispatcher
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

	//Shutdown Bot
	s.Shutdown()
}
