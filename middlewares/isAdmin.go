package middlewares

import (
	"bot/db"
	"bot/helper"
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func IsAdmin(msg *gotgbot.Message) bool {
	tgId := msg.From.Id

	_, err := helper.DB.User.FindFirst(
		db.User.TelegramID.Equals(int(tgId)),
	).Exec(context.Background())

	if err != nil && !IsSudoAdmin(msg) {
		return false
	}

	return true
}

func IsSudoAdmin(msg *gotgbot.Message) bool {
	tgId := msg.From.Id

	return helper.SudoAdmins[tgId]
}
