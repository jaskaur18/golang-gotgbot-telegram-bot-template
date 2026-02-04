package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Environment string

const (
	EnvDev  Environment = "dev"
	EnvQa   Environment = "qa"
	EnvProd Environment = "prod"
)

func (e Environment) String() string {
	return string(e)
}

type Config struct {
	Env      Environment `envconfig:"ENV" default:"dev"`
	Bot      BotConfig
	Database DatabaseConfig
	Logger   LoggerConfig
	Misc     MiscConfig
}

type BotConfig struct {
	Token         string  `envconfig:"BOT_TOKEN" required:"true"`
	SudoAdmins    []int64 `envconfig:"BOT_SUDO_ADMINS"`
	WebhookUrl    string  `envconfig:"WEBHOOK_URL"`
	WebhookSecret string  `envconfig:"WEBHOOK_SECRET"`
}

type DatabaseConfig struct {
	Host             string            `envconfig:"PSQL_HOST" default:"localhost"`
	Port             int               `envconfig:"PSQL_PORT" default:"5432"`
	User             string            `envconfig:"PSQL_USER" default:"postgres"`
	Password         string            `envconfig:"PSQL_PASS" default:"postgres"`
	DBName           string            `envconfig:"PSQL_DBNAME" default:"postgres"`
	SSLMode          string            `envconfig:"PSQL_SSLMODE" default:"disable"`
	MaxOpenConns     int               `envconfig:"PSQL_MAX_OPEN_CONNS" default:"10"`
	MaxIdleConns     int               `envconfig:"PSQL_MAX_IDLE_CONNS" default:"5"`
	ConnMaxLifetime  time.Duration     `envconfig:"PSQL_CONN_MAX_LIFETIME" default:"0s"`
	AdditionalParams map[string]string `envconfig:"PSQL_PARAMS"`
}

type LoggerConfig struct {
	Level              string `envconfig:"LOG_LEVEL" default:"debug"`
	PrettyPrintConsole bool   `envconfig:"LOG_PRETTY" default:"true"`
}

type MiscConfig struct {
	RedisURI   string `envconfig:"REDIS_URI" default:"redis://localhost:6379"`
	LocalesDir string `envconfig:"LOCALES_DIR" default:"locales"`
}

func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// DefaultServiceConfig - kept for backward compatibility during refactor
func DefaultServiceConfig() *Config {
	cfg, err := Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
	return cfg
}

func (c *Config) GetLogLevel() zerolog.Level {
	l, err := zerolog.ParseLevel(c.Logger.Level)
	if err != nil {
		return zerolog.DebugLevel
	}
	return l
}

// ConnectionString generates a PostgreSQL connection string for pgxpool.
func (c *Config) ConnectionString() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.DBName))

	if c.Database.SSLMode != "" {
		b.WriteString(fmt.Sprintf(" sslmode=%s", c.Database.SSLMode))
	} else {
		b.WriteString(" sslmode=disable")
	}

	if c.Database.MaxOpenConns > 0 {
		b.WriteString(fmt.Sprintf(" pool_max_conns=%d", c.Database.MaxOpenConns))
	}
	if c.Database.MaxIdleConns > 0 {
		b.WriteString(fmt.Sprintf(" pool_min_conns=%d", c.Database.MaxIdleConns))
	}
	if c.Database.ConnMaxLifetime > 0 {
		b.WriteString(fmt.Sprintf(" pool_max_conn_lifetime=%s", c.Database.ConnMaxLifetime))
	}

	return b.String()
}
