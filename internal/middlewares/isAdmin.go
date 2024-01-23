package middlewares

import (
	"bot/cmd/bot"
	helper2 "bot/internal/helper"
	"bot/internal/models"
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/volatiletech/null/v8"
)

func IsAdmin(s *bot.Server, msg *gotgbot.Message) bool {
	tgId := msg.From.Id

	if IsSudoAdmin(msg) {
		return true
	}

	u, err := models.Users(models.UserWhere.Telegramid.EQ(null.Int64From(tgId))).One(context.Background(), s.DB)
	if err != nil {
		return false
	}

	if u.Usertype == models.UsertypeADMIN {
		return true
	}

	return false
}

func IsSudoAdmin(msg *gotgbot.Message) bool {
	tgId := msg.From.Id

	return helper2.SudoAdmins[tgId]
}
