package handlers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/middlewares"
	"github.com/rs/zerolog/log"
)

type (
	AccessLevel string
	ChatType    string
)

const (
	AccessLevelSudoAdmin AccessLevel = "sudo admin"
	AccessLevelAdmin                 = "admin"
	AccessLevelUser                  = "user"
	ChatTypePrivate      ChatType    = "private"
	ChatTypeSupergroup               = "supergroup"
	ChatTypeAny                      = "any"
)

func (a AccessLevel) String() string {
	return string(a)
}

func (c ChatType) String() string {
	return string(c)
}

// LoadHandlers loads all message and callback query handlers into the dispatcher.
func LoadHandlers(s *bot.Server) {
	log.Debug().Msg("Loading handlers")
	loadCallbackQueryHandlers(s)
	loadMessageFilterHandler(s)
}

// loadCallbackQueryHandlers loads handlers for processing callback queries.
func loadCallbackQueryHandlers(s *bot.Server) {
	log.Debug().Msg("Loading callback query handlers")
	callbackHandler := handlers.NewCallback(callbackquery.All, handleCallbackQuery(s))
	s.Dispatcher.AddHandler(callbackHandler)
}

func handleCallbackQuery(s *bot.Server) func(b *gotgbot.Bot, c *ext.Context) error {
	return func(b *gotgbot.Bot, c *ext.Context) error {
		if c.CallbackQuery == nil {
			log.Error().Msg("Callback query is nil")
			return errors.New("callback query is nil")
		}
		log.Debug().Str("CallbackData", c.CallbackQuery.Data).Msg("Processing callback query")
		for _, callback := range GetCallbackQueryList() {
			if strings.HasPrefix(c.CallbackQuery.Data, callback.Prefix) {
				if !checkAccessLevel(s, callback.LevelReq, c) {
					return handleNotAllowed(b, c, callback.LevelReq.String(), true)
				}

				if !checkChatType(callback.ChatType, c) {
					return nil
				}

				return callback.Handler(s, b, c)
			}
		}
		return nil
	}
}

func loadMessageFilterHandler(s *bot.Server) {
	msgFilterHandler := handlers.NewMessage(TextMessageFilter, HandleTextMessageFilter(s))
	s.Dispatcher.AddHandler(msgFilterHandler)
}

func TextMessageFilter(msg *gotgbot.Message) bool {
	return msg.Text != ""
}

func HandleTextMessageFilter(s *bot.Server) func(b *gotgbot.Bot, c *ext.Context) error {
	return func(b *gotgbot.Bot, c *ext.Context) error {
		if message.Command(c.EffectiveMessage) {
			return handleCommand(s, b, c)
		}

		log.Debug().Str("Text", c.EffectiveMessage.Text).Msg("Processing text message")

		for _, filter := range GetTextFilters() {
			if filter.Text == c.EffectiveMessage.Text {
				if !checkAccessLevel(s, filter.LevelReq, c) {
					log.Debug().Str("Text", c.EffectiveMessage.Text).Str("LevelReq", filter.LevelReq.String()).Msg("Access denied")
					return handleNotAllowed(b, c, filter.LevelReq.String())
				}

				if !checkChatType(filter.ChatType, c) {
					log.Debug().
						Str("Text", c.EffectiveMessage.Text).
						Str("ChatType", filter.ChatType.String()).
						Msg("Chat type not allowed")
					return nil
				}

				return filter.Handler(s, b, c)
			}
		}

		log.Debug().Str("Text", c.EffectiveMessage.Text).Msg("No handler found for text message")

		return nil
	}
}

func handleCommand(s *bot.Server, b *gotgbot.Bot, c *ext.Context) error {
	text := c.EffectiveMessage.Text

	log.Debug().Str("Command", text).Msg("Processing command")

	for _, cmd := range GetCommandList() {
		if strings.HasPrefix(text, "/"+cmd.Name) {
			if !checkAccessLevel(s, cmd.LevelReq, c) {
				log.Debug().Str("Command", text).Str("LevelReq", cmd.LevelReq.String()).Msg("Access denied")
				return handleNotAllowed(b, c, cmd.LevelReq.String())
			}
			if !checkChatType(cmd.ChatType, c) {
				log.Debug().Str("Command", text).Str("ChatType", cmd.ChatType.String()).Msg("Chat type not allowed")
				return nil
			}
			return cmd.Handler(s, b, c)
		}
	}
	return nil
}

func handleNotAllowed(b *gotgbot.Bot, c *ext.Context, role string, cb ...bool) error {
	var err error
	errMsg := fmt.Sprintf("Access denied: You need %s privileges to use this command", role)
	if len(cb) > 0 && cb[0] {
		_, err = c.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{
			Text: errMsg,
		})
	} else {
		_, err = c.EffectiveMessage.Reply(b, errMsg, nil)
	}

	if err != nil {
		log.Debug().Err(err).Msg("Failed to send access denied message")
	}

	return err
}

func checkAccessLevel(s *bot.Server, levelReq AccessLevel, c *ext.Context) bool {
	switch levelReq {
	case AccessLevelAdmin:
		return middlewares.IsAdmin(s, c.EffectiveMessage)
	case AccessLevelSudoAdmin:
		return middlewares.IsSudoAdmin(c.EffectiveMessage)
	default:
		return true
	}
}

func checkChatType(chatType ChatType, c *ext.Context) bool {
	return chatType == "" || c.EffectiveChat.Type == ChatTypeAny || c.EffectiveChat.Type == chatType.String()
}
