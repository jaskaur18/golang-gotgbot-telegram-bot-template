package helper

import (
	"log"
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
	locales := getLocales(filepath.Join("bots", "storeBot", "locales"))
	loadLocales(filepath.Join("bots", "storeBot", "locales"), locales)
}

func getLocales(root string) []string {
	files, err := os.ReadDir(root)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
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
			log.Fatalf("Error reading locale directory: %v", err)
		}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".ftl") {
				content, err := os.ReadFile(filepath.Join(root, locale, file.Name()))
				if err != nil {
					log.Fatalf("Error reading .ftl file: %v", err)
				}
				resource, errs := fluent.NewResource(string(content))
				if errs != nil {
					log.Printf("Failed to parse .ftl file for locale %s: %v", locale, errs)
					continue
				}
				bundle.AddResourceOverriding(resource)
			}
		}

		bundles[locale] = bundle
	}
}

func GetMessage(locale, key string) (message string) {
	if locale == "" {
		locale = "en"
	}

	bundle, exists := bundles[locale]
	if !exists {
		log.Printf("No translations found for locale: %s", locale)
		return key
	}

	message, _, err := bundle.FormatMessage(key)
	if err != nil {
		log.Printf("Error formatting message: %v", err)
		return key
	}

	return message
}

// CheckLocale Check if Locale Exists
func CheckLocale(locale string) bool {
	_, exists := bundles[locale]
	return exists
}
