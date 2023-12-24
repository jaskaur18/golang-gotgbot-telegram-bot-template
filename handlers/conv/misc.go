package conv

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"strings"
)

func literalyAnything(_ *gotgbot.Message) bool {
	return true
}

func noCommands(msg *gotgbot.Message) bool {
	return message.Text(msg) && !message.Command(msg)
}

func CancelMsg(msg *gotgbot.Message) bool {
	return strings.ToLower(msg.Text) == "cancel" || strings.ToLower(msg.Text) == "/cancel"
}

func HandleCancel(b *gotgbot.Bot, ctx *ext.Context) error {
	_, _ = ctx.EffectiveMessage.Reply(b, "Cancelled", nil)
	return handlers.EndConversation()
}
