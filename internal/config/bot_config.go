package config

import (
	"context"

	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/pkl/pklgen"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/pkl/pklgen/environment"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Bot struct {
	TelegramBotConfig *pklgen.BOT
	Database          *pklgen.Database
	Logger            *pklgen.LOGGER
	Misc              *pklgen.MISC
	ENV               environment.Environment
}

func DefaultServiceConfig() Bot {
	cfg, err := pklgen.LoadFromPath(context.Background(), "pkl/local/BotConfig.pkl")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	return Bot{
		TelegramBotConfig: cfg.Bot,
		Database:          cfg.DB,
		Logger:            cfg.Logger,
		Misc:              cfg.Misc,
		ENV:               cfg.ENV,
	}
}

func LogLevelFromString(s string) zerolog.Level {
	l, err := zerolog.ParseLevel(s)
	if err != nil || l == zerolog.NoLevel {
		log.Error().Err(err).Msgf("Failed to parse log level, defaulting to %s", zerolog.DebugLevel)
		return zerolog.DebugLevel
	}

	return l
}
