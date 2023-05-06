package commands

import (
	"bot/model"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"strings"
	"time"
)

func CommandBroadcast(b *gotgbot.Bot, c *ext.Context) error {
	users, err := model.GetAllUsers()
	if err != nil {
		_, err := c.Message.Reply(b, "Error Happened", nil)
		if err != nil {
			return err
		}
		return err
	}

	msg := strings.Join(c.Args()[1:], " ")
	if msg == "" {
		c.Message.Reply(b, "Please enter a message", nil)
		return nil
	}

	c.Message.Reply(b, "Broadcasting...", nil)

	totalSend := 0
	for _, user := range users {
		_, err = b.SendMessage(user.TelegramId, msg, nil)
		if err != nil {
			c.Message.Reply(b, "Error Happened", nil)
			return err
		}
		totalSend++
		time.Sleep(100 * time.Millisecond)
	}

	msg = fmt.Sprintf("Broadcasted to %d users", totalSend)
	c.Message.Reply(b, msg, nil)

	return nil
}
