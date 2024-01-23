package helper

import (
	"bot/internal/config"
	"time"
)

var SudoAdmins = map[int64]bool{}
var RedisTimeOut = 10 * time.Second

func InitConstants(c *config.Bot) {
	sIDs := c.GetAdmins()
	for _, id := range sIDs {
		SudoAdmins[id] = true
	}
}
