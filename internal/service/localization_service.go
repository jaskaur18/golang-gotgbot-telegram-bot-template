package service

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/domain"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/repository/redis"
	"github.com/lus/fluent.go/fluent"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
)

type LocalizationService struct {
	bundles     map[string]*fluent.Bundle
	sessionRepo *redis.SessionRepository
}

func NewLocalizationService(localesDir string, sessionRepo *redis.SessionRepository) *LocalizationService {
	l := &LocalizationService{
		bundles:     make(map[string]*fluent.Bundle),
		sessionRepo: sessionRepo,
	}
	l.loadLocales(localesDir)
	return l
}

func (s *LocalizationService) loadLocales(root string) {
	files, err := os.ReadDir(root)
	if err != nil {
		log.Error().Err(err).Msg("Error reading locale directory, ensuring directory exists")
		return
	}

	for _, f := range files {
		if f.IsDir() {
			locale := f.Name()
			bundle := fluent.NewBundle(language.Make(locale))

			localeFiles, err := os.ReadDir(filepath.Join(root, locale))
			if err != nil {
				continue
			}

			for _, lf := range localeFiles {
				if strings.HasSuffix(lf.Name(), ".ftl") {
					content, err := os.ReadFile(filepath.Join(root, locale, lf.Name()))
					if err != nil {
						continue
					}
					resource, errs := fluent.NewResource(string(content))
					if len(errs) > 0 {
						for _, e := range errs {
							log.Error().Err(e).Str("file", lf.Name()).Msg("Failed to parse FTL")
						}
					}
					bundle.AddResourceOverriding(resource)
				}
			}
			s.bundles[locale] = bundle
		}
	}
}

func (s *LocalizationService) GetAvailableLocales() []string {
	var locales []string
	for k := range s.bundles {
		locales = append(locales, k)
	}
	return locales
}

func (s *LocalizationService) CheckLocale(locale string) bool {
	_, ok := s.bundles[locale]
	return ok
}

func (s *LocalizationService) SetUserLocale(ctx context.Context, telegramID int64, locale string) error {
	session, err := s.sessionRepo.Get(ctx, telegramID)
	if err != nil {
		// If get fails, maybe create? but repo handles default return if Nil.
		// If real error, fail.
		return err
	}
	// If nil wrapping, session repo Get returns default struct if not found.

	// Create/Update session
	if session == nil {
		session = &domain.Session{TelegramID: telegramID}
	}
	session.Language = locale
	return s.sessionRepo.Save(ctx, session)
}

func (s *LocalizationService) GetUserLocale(ctx context.Context, telegramID int64) string {
	session, err := s.sessionRepo.Get(ctx, telegramID)
	if err != nil || session == nil {
		return "en"
	}
	return session.Language
}

func (s *LocalizationService) GetMessage(ctx context.Context, telegramID int64, key string) string {
	return s.GetMessageWithArgs(ctx, telegramID, key)
}

// GetMessageWithArgs renamed from Options to be clearer or just fix signature
func (s *LocalizationService) GetMessageWithOptions(ctx context.Context, telegramID int64, key string, opts ...*fluent.FormatContext) string {
	locale := s.GetUserLocale(ctx, telegramID)

	bundle, ok := s.bundles[locale]
	if !ok {
		bundle = s.bundles["en"]
	}
	if bundle == nil {
		return key
	}

	msg, _, err := bundle.FormatMessage(key, opts...)
	if err != nil {
		return key
	}
	return msg
}

// Alias for simpler usage
func (s *LocalizationService) GetMessageWithArgs(ctx context.Context, telegramID int64, key string, opts ...*fluent.FormatContext) string {
	return s.GetMessageWithOptions(ctx, telegramID, key, opts...)
}
