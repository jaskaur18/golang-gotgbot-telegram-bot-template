package helper

import (
	"errors"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type env struct {
	BotToken      string `validate:"required" json:"STORE_BOT_TOKEN"`
	SudoAdmins    string `validate:"required" json:"SUDO_ADMINS"`
	RedisUri      string `validate:"required" json:"REDIS_URI"`
	PROD          bool   `validate:"boolean" json:"PROD"`
	WebhookUrl    string `validate:"required_if=PROD true" json:"WEBHOOK_URL"`
	WebhookSecret string `validate:"required_if=PROD true" json:"WEBHOOK_SECRET"`
}

var Env env

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}

	envType := reflect.TypeOf(Env)
	envValue := reflect.ValueOf(&Env).Elem()

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		envVarName := pair[0]
		envVarValue := pair[1]

		for i := 0; i < envType.NumField(); i++ {
			field := envType.Field(i)

			jsonTag := field.Tag.Get("json")
			if jsonTag == envVarName {
				if field.Type.Name() == "string" {
					envValue.FieldByName(field.Name).SetString(envVarValue)
				} else if field.Type.Name() == "bool" {
					b, err := strconv.ParseBool(envVarValue)
					if err != nil {
						log.Fatal("Error parsing boolean value from environment variable: ", err)
					}
					envValue.FieldByName(field.Name).SetBool(b)
				}
			}
		}
	}

	validate := validator.New()

	err = validate.Struct(Env)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			log.Fatal("Error validating environment variables: ", err)
			return
		}

		log.Fatal("Error validating environment variables: ", err)
		return
	}

	log.Printf("Environment variables loaded successfully 🚀\n%+v\n", Env)
}
