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

<<<<<<< HEAD
	//_, err := helper.DB.User.FindFirst(
	//	db.User.TelegramID.Equals(db.BigInt(tgId)),
	//).Exec(context.Background())
	//
	//if err != nil {
	//	return false
	//}
=======
	_, err := helper.DB.User.FindFirst(
		db.User.TelegramID.Equals(int(tgId)),
	).Exec(context.Background())
>>>>>>> parent of dc24b0d (Update i18n implementation, libraries and installation script)

	return false
}

func IsSudoAdmin(msg *gotgbot.Message) bool {
	tgId := msg.From.Id

	return helper.SudoAdmins[tgId]
}
