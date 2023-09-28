package middlewares

import (
	"bot/db"
	"bot/helper"
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func IsAdmin(msg *gotgbot.Message) bool {
	tgId := msg.From.Id

	if IsSudoAdmin(msg) {
		return true
	}

	_, err := helper.DB.User.FindFirst(
		db.User.TelegramID.Equals(db.BigInt(tgId)),
	).Exec(context.Background())

	if err != nil {
		return false
	}

	return true
}

func IsSudoAdmin(msg *gotgbot.Message) bool {
	tgId := msg.From.Id

	return helper.SudoAdmins[tgId]
}
