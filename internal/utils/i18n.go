package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"github.com/lus/fluent.go/fluent"
	"golang.org/x/text/language"
)

type LocaleLoader struct {
	bundles map[string]*fluent.Bundle
}

func NewLocaleLoader(localesDir string) *LocaleLoader {
	loader := &LocaleLoader{
		bundles: make(map[string]*fluent.Bundle),
	}
	locales := loader.getLocales(localesDir)
	loader.loadLocales(localesDir, locales)
	return loader
}

func (l *LocaleLoader) getLocales(root string) []string {
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

func (l *LocaleLoader) loadLocales(root string, locales []string) {
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
					for _, err := range errs {
						log.Error().Err(err).Msg("Failed to parse .ftl file")
					}
					continue
				}
				bundle.AddResourceOverriding(resource)
			}
		}

		l.bundles[locale] = bundle
	}
}

func (l *LocaleLoader) GetUserLocale(s *redis.Client, tgID int64) string {
	locale := "en"

	session, err := GetSession(s, tgID) // Assuming GetSession is implemented elsewhere
	if err != nil {
		log.Debug().
			Int64("tgId", tgID).
			Err(err).Msg("Error getting session")
		return locale
	}

	if session.Language != "" {
		locale = session.Language
	}
	return locale
}

func (l *LocaleLoader) GetMessage(
	r *redis.Client, c *ext.Context, key string, contexts ...*fluent.FormatContext) string {
	tgID := c.EffectiveUser.Id
	locale := l.GetUserLocale(r, tgID)

	bundle, exists := l.bundles[locale]
	if !exists {
		log.Warn().
			Int64("tgId", tgID).
			Str("locale", locale).
			Msg("No translations found for locale")
		return key
	}

	message, _, err := bundle.FormatMessage(key, contexts...)
	if err != nil {
		log.Info().
			Int64("tgId", tgID).
			Str("locale", locale).
			Err(err).Msg("Error formatting message")
		return key
	}

	return message
}

func (l *LocaleLoader) CheckLocale(locale string) bool {
	_, exists := l.bundles[locale]
	return exists
}
