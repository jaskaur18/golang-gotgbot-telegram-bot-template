package bot

import (
	"errors"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/config"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/db"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/utils"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/pkl/pklgen/environment"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Config     config.Bot
	Bot        *gotgbot.Bot
	Updater    *ext.Updater
	Dispatcher *ext.Dispatcher
	DB         *pgxpool.Pool
	Queries    *db.Queries
	Redis      *redis.Client
}

func NewServer(config config.Bot, db *pgxpool.Pool, queries *db.Queries, redisClient *redis.Client) *Server {
	b := &Server{
		Config:  config,
		DB:      db,
		Queries: queries,
		Redis:   redisClient,
	}

	return b
}

func (s *Server) Ready() bool {
	return s.DB != nil &&
		s.Bot != nil
}

func (s *Server) Start() error {
	if !s.Ready() {
		return errors.New("bot is not ready")
	}

	var err error

	if s.Config.ENV == environment.Prod {
		err = utils.ProdLaunch(s.Bot, s.Updater)
	} else {
		err = utils.DevLaunch(s.Bot, s.Updater)
	}

	return err
}

func (s *Server) Shutdown() {
	log.Warn().Msg("Shutting down server")

	if s.DB != nil {
		log.Debug().Msg("Closing database connection")

		s.DB.Close()
	}

	log.Debug().Msg("Shutting down echo server")

	s.Updater.StopAllBots()
}
