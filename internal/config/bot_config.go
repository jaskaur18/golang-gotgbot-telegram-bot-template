package config

import (
	"bot/internal/env"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
	"time"
)

type LoggerServer struct {
	Level              zerolog.Level
	RequestLevel       zerolog.Level
	LogRequestBody     bool
	LogRequestHeader   bool
	LogRequestQuery    bool
	LogResponseBody    bool
	LogResponseHeader  bool
	LogCaller          bool
	PrettyPrintConsole bool
}

type TelegramBotConfig struct {
	BotToken   string
	SudoAdmins string
	Prod       bool
}

type ManagementServer struct {
	WebhookUrl    string
	WebhookSecret string
	RedisUri      string
}

type Bot struct {
	TelegramBotConfig TelegramBotConfig
	Database          Database
	Logger            LoggerServer
	ManagementServer  ManagementServer
}

func DefaultServiceConfigFromEnv() Bot {
	//Init All The Env Variables
	env.InitEnv()

	return Bot{
		TelegramBotConfig: TelegramBotConfig{
			BotToken:   env.Config.BotToken,
			SudoAdmins: env.Config.SudoAdmins,
			Prod:       env.Config.PROD,
		},
		Database: Database{
			Host:             env.Config.PSQLHOST,
			Port:             env.Config.PSQLPORT,
			Username:         env.Config.PSQLUSER,
			Password:         env.Config.PSQLPASS,
			Database:         env.Config.PSQLDB,
			AdditionalParams: map[string]string{"sslmode": env.Config.PSQLSSLMODE},
			MaxOpenConns:     env.Config.DBMaxOpenConns,
			MaxIdleConns:     env.Config.MaxIdleConns,
			ConnMaxLifetime:  time.Duration(env.Config.ConnectionMaxLifetime) * time.Second,
		},
		Logger: LoggerServer{
			Level:              LogLevelFromString(env.Config.LoggerLevel),
			RequestLevel:       LogLevelFromString(env.Config.LoggerRequestLevel),
			LogRequestBody:     env.Config.LoggerLogRequestBody,
			LogRequestHeader:   env.Config.LoggerLogRequestHeader,
			LogRequestQuery:    env.Config.LoggerLogRequestQuery,
			LogResponseBody:    env.Config.LoggerLogResponseBody,
			LogResponseHeader:  env.Config.LoggerLogResponseHeader,
			LogCaller:          env.Config.LoggerLogCaller,
			PrettyPrintConsole: env.Config.LoggerPrettyPrintConsole,
		},
		ManagementServer: ManagementServer{
			WebhookUrl:    env.Config.WebhookUrl,
			WebhookSecret: env.Config.WebhookSecret,
			RedisUri:      env.Config.RedisURI,
		},
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

func (b *Bot) GetAdmins() []int64 {
	admins := make([]int64, 0)
	for _, admin := range strings.Split(b.TelegramBotConfig.SudoAdmins, ",") {
		adminInt, err := strconv.ParseInt(admin, 10, 64)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse admin id")
			continue
		}
		admins = append(admins, adminInt)
	}
	return admins
}
