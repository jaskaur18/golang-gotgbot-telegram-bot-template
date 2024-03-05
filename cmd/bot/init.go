package bot

import (
	"context"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/config"

	// lint:ignore ST1001 This is required to initialize the database
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func (s *Server) InitBot() error {
	bot, err := gotgbot.NewBot(s.Config.TelegramBotConfig.Token, &gotgbot.BotOpts{})

	if err != nil {
		return err
	}

	s.Bot = bot

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If a handler returns an error, log it and continue going.
		Error: func(_ *gotgbot.Bot, _ *ext.Context, err error) ext.DispatcherAction {
			log.Error().Stack().Err(errors.Wrap(err, "wrapped error")).Msg("an error occurred while handling update")
			return ext.DispatcherActionNoop
		},
		Panic: func(_ *gotgbot.Bot, _ *ext.Context, r interface{}) {
			log.Error().Stack().Any("panic_reason", r).Msg("a panic occurred while handling update")
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	// Create updater and dispatcher.
	updater := ext.NewUpdater(dispatcher, &ext.UpdaterOpts{})

	s.Updater = updater
	s.Dispatcher = dispatcher

	return nil
}

func InitDB(c *config.Bot) (*pgxpool.Pool, error) {
	pgxPoolConfig, err := pgxpool.ParseConfig(c.ConnectionString())
	if err != nil {
		log.Error().Err(err).Msg("Error parsing pgxpool config")
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(context.TODO(), pgxPoolConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func InitRedis(ctx context.Context, redisURI string) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Err(err).Str("URI", redisURI).Msg("Error parsing redis url")
		return nil, err
	}

	r := redis.NewClient(opt)

	_, err = r.Ping(ctx).Result()

	if err != nil {
		log.Err(err).Msg("Error connecting to Redis")
		return nil, err
	}

	return r, nil
}
