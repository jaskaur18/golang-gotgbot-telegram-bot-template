package helpers

import (
	"strings"
)

var AdminsIds = []int64{0, 0}

func InitConstants() {
	adminid := strings.Split(Env.AdminIds, ",")
	for _, id := range adminid {
		AdminIdInt, ok := StringToInt64(id)
		if !ok {
			continue
		}
		AdminsIds = append(AdminsIds, AdminIdInt)
	}
}
