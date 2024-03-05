package constant

import (
	"time"

	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/config"
)

var sudoAdmins = map[int64]bool{} //nolint:gochecknoglobals

func GetSudoAdmins() map[int64]bool {
	return sudoAdmins
}

func GetRedisTimeOut() time.Duration {
	var redisTimeOut = 10
	return time.Duration(redisTimeOut) * time.Second
}

func InitConstants(c *config.Bot) {
	sIDs := c.TelegramBotConfig.SudoAdmins
	for _, id := range sIDs {
		sudoAdmins[int64(id)] = true
	}
}
