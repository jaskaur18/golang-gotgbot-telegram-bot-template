package middlewares

import (
	"context"
	"encoding/json"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type i18nClient struct {
	gotgbot.BotClient
}

func (b i18nClient) RequestWithContext(ctx context.Context, token string, method string, params map[string]string, data map[string]gotgbot.NamedReader, opts *gotgbot.RequestOpts) (json.RawMessage, error) {

	// Handle local parse mode
	if strings.Contains(params["parse_mode"], "local") {
		parseModes := strings.Split(params["parse_mode"], ",")
		var cleanedParseModes []string
		for _, parseMode := range parseModes {
			if parseMode != "local" {
				cleanedParseModes = append(cleanedParseModes, parseMode)
			}
		}
		// Set parse_mode to the first valid mode or empty if none left
		params["parse_mode"] = ""
		if len(cleanedParseModes) > 0 {
			params["parse_mode"] = cleanedParseModes[0]
		}

		chatId, err := strconv.ParseInt(params["chat_id"], 10, 64)
		if err != nil {
			log.Printf("Error while parsing ChatId, chat_id: %s, error: %v", params["chat_id"], err)
			return nil, err
		}

		s, err := helper.GetSession(chatId)
		if err != nil {
			log.Printf("Error while getting session, user_id: %d, error: %v", chatId, err)
			return nil, err
		}

		msg := helper.GetMessage(s.Language, params["text"])
		params["text"] = msg
	}

	val, err := b.BotClient.RequestWithContext(ctx, token, method, params, data, opts)
	if err != nil {
		log.Printf("Warning, got an error: %v", err)
	}
	return val, err
}

func NewI18nClient() gotgbot.BotClient {
	return i18nClient{
		BotClient: &gotgbot.BaseBotClient{
			Client:             http.Client{},
			UseTestEnvironment: false,
			DefaultRequestOpts: &gotgbot.RequestOpts{
				Timeout: gotgbot.DefaultTimeout,
				APIURL:  gotgbot.DefaultAPIURL,
			},
		},
	}
}
