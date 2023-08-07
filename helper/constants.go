package helper

import (
	"strings"
)

var SudoAdmins = map[int64]bool{}

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
