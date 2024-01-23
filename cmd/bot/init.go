package bot

import (
	"bot/internal/config"
	"context"
	"database/sql"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func (s *Server) InitBot() error {
	bot, err := gotgbot.NewBot(s.Config.TelegramBotConfig.BotToken, &gotgbot.BotOpts{})

	if err != nil {
		return err
	}

	s.Bot = bot

	// Create updater and dispatcher.
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			// If a handler returns an error, log it and continue going.
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Error().Stack().Err(errors.Wrap(err, "wrapped error")).Msg("an error occurred while handling update")
				return ext.DispatcherActionNoop
			},
			Panic: func(b *gotgbot.Bot, ctx *ext.Context, r interface{}) {
				log.Error().Stack().Any("panic_reason", r).Msg("a panic occurred while handling update")
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		}),
	})

	s.Updater = updater
	s.Dispatcher = updater.Dispatcher

	return nil
}

func InitDB(ctx context.Context, d *config.Database) (*sql.DB, error) {
	log.Printf(d.ConnectionString())
	db, err := sql.Open("postgres", d.ConnectionString())
	if err != nil {
		return nil, err
	}

	if d.MaxOpenConns > 0 {
		db.SetMaxOpenConns(d.MaxOpenConns)
	}
	if d.MaxIdleConns > 0 {
		db.SetMaxIdleConns(d.MaxIdleConns)
	}
	if d.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(d.ConnMaxLifetime)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func InitRedis(ctx context.Context, redisUri string) (*redis.Client, error) {
	opt, err := redis.ParseURL(redisUri)
	if err != nil {
		log.Err(err).Str("URI", redisUri).Msg("Error parsing redis url")
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
