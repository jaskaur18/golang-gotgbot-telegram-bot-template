package helper

import (
	"strings"
	"time"
)

var SudoAdmins = map[int64]bool{}
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
