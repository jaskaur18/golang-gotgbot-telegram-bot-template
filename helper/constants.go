package helper

import (
	"strings"
	"time"
)

var SupportGuyUsername = "xxx"
var AdminUsername = "xxx"
var SupportBotUsername = "xxx"

var SudoAdmins = map[int64]bool{}
var BanUsers = map[int64]bool{}
var RedisTimeOut = 10 * time.Second

func InitConstants() {
	sIDs := strings.Split(Env.SudoAdmins, ",")
	for _, id := range sIDs {
		SudoAdminInt, ok := StringToInt64(id)
		if !ok {
			continue
		}
		SudoAdmins[SudoAdminInt] = true
	}
}
