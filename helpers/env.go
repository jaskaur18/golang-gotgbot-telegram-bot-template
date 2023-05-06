package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type env struct {
	BotToken      string `validate:"required"`
	PostgresUri   string `validate:"required"`
	AdminIds      string `validate:"required"`
	PROD          bool   `validate:"boolean"`
	WebhookUrl    string `validate:"required_if=PROD true"`
	WebhookSecret string `validate:"required_if=PROD true"`
}

var Env *env

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file: ", err)
	}

	Env = &env{
		BotToken:      os.Getenv("BOT_TOKEN"),
		PostgresUri:   os.Getenv("POSTGRES_URI"),
		AdminIds:      os.Getenv("ADMIN_IDS"),
		PROD:          os.Getenv("PROD") == "true",
		WebhookUrl:    os.Getenv("WEBHOOK_URL"),
		WebhookSecret: os.Getenv("WEBHOOK_SECRET"),
	}

	validate := validator.New()

	err = validate.Struct(Env)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Fatal("Error validating environment variables: ", err)
			return
		}

		log.Fatal("Error validating environment variables: ", err)
		return
	}

	log.Printf("Environment variables loaded successfully 🚀\n%+v\n", Env)
}
