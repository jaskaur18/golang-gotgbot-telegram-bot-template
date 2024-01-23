package utils

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lus/fluent.go/fluent"
	"golang.org/x/text/language"
)

// Global map to store bundles for each locale
var bundles = make(map[string]*fluent.Bundle)

func init() {
	// Load locales from the locales directory
	locales := getLocales("locales")
	loadLocales("locales", locales)
}

func getLocales(root string) []string {
	files, err := os.ReadDir(root)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading locale directory")
	}

	var locales []string
	for _, f := range files {
		if f.IsDir() {
			locales = append(locales, f.Name())
		}
	}

	return locales
}

func loadLocales(root string, locales []string) {
	for _, locale := range locales {
		bundle := fluent.NewBundle(language.Make(locale))
		files, err := os.ReadDir(filepath.Join(root, locale))
		if err != nil {
			log.Fatal().Err(err).Msg("Error reading locale directory")
		}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".ftl") {
				content, err := os.ReadFile(filepath.Join(root, locale, file.Name()))
				if err != nil {
					log.Err(err).Msg("Error reading .ftl file")
					continue
				}
				resource, errs := fluent.NewResource(string(content))
				if errs != nil {
					log.Error().Any("errors", errs).Msg("Failed to parse .ftl file")
					continue
				}
				bundle.AddResourceOverriding(resource)
			}
		}

		bundles[locale] = bundle
	}
}

func GetUserLocale(s *redis.Client, tgId int64) (locale string) {
	locale = "en"

	session, err := GetSession(s, tgId)
	if err != nil {
		log.Debug().
			Int64("tgId", tgId).
			Err(err).Msg("Error getting session")
		return locale
	}

	if session.Language != "" {
		return session.Language
	}

	return locale
}

func GetMessage(r *redis.Client, c *ext.Context, key string, contexts ...*fluent.FormatContext) (message string) {
	tgId := c.EffectiveUser.Id
	locale := GetUserLocale(r, tgId)

	bundle, exists := bundles[locale]
	if !exists {
		log.Warn().
			Int64("tgId", tgId).
			Str("locale", locale).
			Msg("No translations found for locale")

		return key
	}

	message, _, err := bundle.FormatMessage(key, contexts...)
	if err != nil {
		log.Info().
			Int64("tgId", tgId).
			Str("locale", locale).
			Err(err).Msg("Error formatting message")
		return key
	}

	return message
}

// CheckLocale Check if Locale Exists
func CheckLocale(locale string) bool {
	_, exists := bundles[locale]
	return exists
}
