package middlewares

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
)

func IsAdmin(msg *gotgbot.Message) bool {
	//tgId := msg.From.Id

	if IsSudoAdmin(msg) {
		return true
	}

	//_, err := helper.DB.User.FindFirst(
	//	db.User.TelegramID.Equals(db.BigInt(tgId)),
	//).Exec(context.Background())
	//
	//if err != nil {
	//	return false
	//}

	return false
}

func IsSudoAdmin(msg *gotgbot.Message) bool {
	tgId := msg.From.Id

	return helper.SudoAdmins[tgId]
}
