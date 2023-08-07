package commands

import (
	"bot/handlers/misc"
	"bot/helper"
	"context"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
	"strings"
	"time"
)

func CommandBroadcast(b *gotgbot.Bot, c *ext.Context) error {

	allUsers, err := helper.DB.User.FindMany().Exec(context.Background())
	if err != nil {
		return misc.ErrorHandler(b, c, err)
	}

	msg := strings.Join(c.Args()[1:], " ")
	if msg == "" {
		c.Message.Reply(b, "Please enter a message", nil)
		return nil
	}

	_, _ = c.Message.Reply(b, "Broadcasting...", nil)

	totalSend := 0
	for _, user := range allUsers {
		_, err = b.SendMessage(int64(user.TelegramID), msg, nil)
		if err != nil {
			log.Printf("Error sending message to %d: %s", user.TelegramID, err)
			continue
		}
		totalSend++
		time.Sleep(100 * time.Millisecond)
	}

	msg = fmt.Sprintf("Broadcasted to %d users", totalSend)
	_, _ = c.Message.Reply(b, msg, nil)

	return nil
}
