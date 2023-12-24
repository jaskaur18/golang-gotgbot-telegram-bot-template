package commands

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers/misc"
	"github.com/jaskaur18/moimoiStoreBot/helpers"
	"github.com/pocketbase/dbx"
	"log"
	"strings"
	"time"
)

func CommandBroadcast(b *gotgbot.Bot, c *ext.Context) error {

	allUsers, err := helpers.PB.Dao().FindRecordsByFilter("tgUser", "", "created", 0, 0, dbx.Params{})
	if err != nil {
		return misc.ErrorHandler(b, c, err)
	}

	msg := strings.Join(c.Args()[1:], " ")
	if msg == "" {
		_, _ = c.EffectiveMessage.Reply(b, "Please enter a message", nil)
		return nil
	}

	_, _ = c.Message.Reply(b, "Broadcasting...", nil)

	totalSend := 0
	for _, user := range allUsers {
		tgID := user.Get("tgId").(int64)
		_, err = b.SendMessage(tgID, msg, nil)
		if err != nil {
			log.Printf("Error sending message to %d: %s", tgID, err)
			continue
		}
		totalSend++
		time.Sleep(100 * time.Millisecond)
	}

	msg = fmt.Sprintf("Broadcasted to %d users", totalSend)
	_, _ = c.Message.Reply(b, msg, nil)

	return nil
}
