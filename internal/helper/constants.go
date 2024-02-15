package helper

import (
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/config"
	"time"
)

var SudoAdmins = map[int64]bool{}
var RedisTimeOut = 10 * time.Second

func InitConstants(c *config.Bot) {
	sIDs := c.TelegramBotConfig.SudoAdmins
	for _, id := range sIDs {
		SudoAdmins[int64(id)] = true
	}
}
