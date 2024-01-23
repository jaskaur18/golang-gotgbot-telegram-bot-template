package env

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"log"
)

type Env struct {
	BotToken                 string `env:"BOT_TOKEN,required"`
	PSQLDB                   string `env:"PSQL_DBNAME,required"`
	PSQLHOST                 string `env:"PSQL_HOST,required"`
	PSQLPORT                 int    `env:"PSQL_PORT,default=5432"`
	PSQLUSER                 string `env:"PSQL_USER,required"`
	PSQLPASS                 string `env:"PSQL_PASS,required"`
	PSQLSSLMODE              string `env:"PSQL_SSLMODE,default=disable"`
	RedisURI                 string `env:"REDIS_URI,required"`
	DBMaxOpenConns           int    `env:"DB_MAX_OPEN_CONNS,default=25"`
	MaxIdleConns             int    `env:"DB_MAX_IDLE_CONNS,default=1"`
	ConnectionMaxLifetime    int    `env:"DB_CONN_MAX_LIFETIME,default=14400"`
	SudoAdmins               string `env:"SUDO_ADMINS,required"`
	PROD                     bool   `env:"PROD,default=false"`
	WebhookUrl               string `env:"WEBHOOK_URL"`
	WebhookSecret            string `env:"WEBHOOK_SECRET"`
	LoggerLevel              string `env:"LOGGER_LEVEL,default=info"`
	LoggerRequestLevel       string `env:"LOGGER_REQUEST_LEVEL,default=info"`
	LoggerLogRequestBody     bool   `env:"LOGGER_LOG_REQUEST_BODY,default=false"`
	LoggerLogRequestHeader   bool   `env:"LOGGER_LOG_REQUEST_HEADER,default=false"`
	LoggerLogRequestQuery    bool   `env:"LOGGER_LOG_REQUEST_QUERY,default=false"`
	LoggerLogResponseBody    bool   `env:"LOGGER_LOG_RESPONSE_BODY,default=false"`
	LoggerLogResponseHeader  bool   `env:"LOGGER_LOG_RESPONSE_HEADER,default=false"`
	LoggerLogCaller          bool   `env:"LOGGER_LOG_CALLER,default=false"`
	LoggerPrettyPrintConsole bool   `env:"LOGGER_PRETTY_PRINT_CONSOLE,default=false"`
}

var Config Env

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}

	ctx := context.Background()

	if err := envconfig.Process(ctx, &Config); err != nil {
		log.Fatalf("Error processing environment variables: %v", err)
	}

	log.Printf("Environment variables loaded successfully 🚀\n%+v\n", Config)
}
