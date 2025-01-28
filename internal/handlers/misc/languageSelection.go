package misc

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/utils"
)

func HandleLanguageInline(s *bot.Server, b *gotgbot.Bot, ctx *ext.Context) error {
	// Fetch available locales
	availableLocales := s.Locale.GetAvailableLocales()
	curentLang := s.Locale.GetUserLocale(s.Redis, ctx.EffectiveUser.Id)

	// Build the inline keyboard dynamically
	var keyboard [][]gotgbot.InlineKeyboardButton
	for _, locale := range availableLocales {
		flag := getCountryFlag(locale) // Get the flag for the locale
		if locale == curentLang {
			flag = "✅ " + flag // Add a checkmark if the locale is the current language
		}

		buttonText := fmt.Sprintf("%s %s", flag, strings.ToUpper(locale))
		keyboard = append(keyboard, []gotgbot.InlineKeyboardButton{
			{
				Text:         buttonText,
				CallbackData: fmt.Sprintf("setLang:%s", locale),
			},
		})
	}

	// Send the inline keyboard to the user
	_, err := ctx.Message.Reply(b, s.Locale.GetMessage(s.Redis, ctx, "langChoose"), &gotgbot.SendMessageOpts{
		ReplyMarkup: &gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		},
		ParseMode: "HTML",
	})
	return err
}

func HandleSetLanguageCallback(s *bot.Server, b *gotgbot.Bot, ctx *ext.Context) error {
	data := ctx.CallbackQuery.Data
	if !strings.HasPrefix(data, "setLang:") {
		return nil
	}

	locale := strings.TrimPrefix(data, "setLang:")

	if !s.Locale.CheckLocale(locale) {
		_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      s.Locale.GetMessage(s.Redis, ctx, "invalidLang"),
			ShowAlert: true,
		})
		return err
	}

	session, err := utils.GetSession(s.Redis, ctx.EffectiveUser.Id)
	if err != nil {
		_, err := ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text:      s.Locale.GetMessage(s.Redis, ctx, "saveState"),
			ShowAlert: true,
		})
		return err
	}

	session.Language = locale
	err = session.Save()
	if err != nil {
		return err
	}

	_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
		Text: s.Locale.GetMessage(s.Redis, ctx, "langSuccess"),
	})

	return err
}

// Helper function to get a country flag emoji based on ISO 639-1 language code
func getCountryFlag(locale string) string {
	// Map ISO 639-1 language codes to ISO 3166-1 alpha-2 country codes
	langToCountry := map[string]string{
		"en": "US", // English -> United States
		"es": "ES", // Spanish -> Spain
		"fr": "FR", // French -> France
		"de": "DE", // German -> Germany
		"it": "IT", // Italian -> Italy
		"ru": "RU", // Russian -> Russia
		"ja": "JP", // Japanese -> Japan
		"zh": "CN", // Chinese -> China
		"ar": "SA", // Arabic -> Saudi Arabia
		"pt": "PT", // Portuguese -> Portugal
		"hi": "IN", // Hindi -> India
	}

	countryCode, exists := langToCountry[locale]
	if !exists {
		return "❓" // Return a question mark if the mapping doesn't exist
	}

	// Convert country code to flag emoji
	flag := ""
	for _, char := range countryCode {
		flag += string(rune(char) + 0x1F1A5) // Offset for regional indicator symbols
	}
	return flag
}
