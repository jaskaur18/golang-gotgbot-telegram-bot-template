package bot

import (
	"bot/internal/config"
	"bot/internal/utils"
	"database/sql"
	"errors"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Config     config.Bot
	Bot        *gotgbot.Bot
	Updater    *ext.Updater
	Dispatcher *ext.Dispatcher
	DB         *sql.DB
	Redis      *redis.Client
}

func NewServer(config config.Bot, db *sql.DB, redisClient *redis.Client) *Server {
	b := &Server{
		Config: config,
		DB:     db,
		Redis:  redisClient,
	}

	return b
}

func (b *Server) Ready() bool {
	return b.DB != nil &&
		b.Bot != nil
}

func (b *Server) Start() error {
	if !b.Ready() {
		return errors.New("bot is not ready")
	}

	var err error

	if b.Config.TelegramBotConfig.Prod {
		err = utils.ProdLaunch(b.Bot, b.Updater)
	} else {
		err = utils.DevLaunch(b.Bot, b.Updater)
	}

	return err
}

func (b *Server) Shutdown() {
	log.Warn().Msg("Shutting down server")

	if b.DB != nil {
		log.Debug().Msg("Closing database connection")

		if err := b.DB.Close(); err != nil && !errors.Is(err, sql.ErrConnDone) {
			log.Error().Err(err).Msg("Failed to close database connection")
		}
	}

	log.Debug().Msg("Shutting down echo server")

	b.Updater.StopAllBots()
}
