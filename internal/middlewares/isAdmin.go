package middlewares

import (
	"context"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/db"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/helper"
)

func IsAdmin(s *bot.Server, msg *gotgbot.Message) bool {
	tgId := msg.From.Id

	if IsSudoAdmin(msg) {
		return true
	}

	u, err := s.Queries.GetUserByTelegramID(context.Background(), pgtype.Int8{Int64: tgId, Valid: true})
	if err != nil {
		return false
	}

	if u.UserType == db.UsertypeADMIN {
		return true
	}

	return false
}

func IsSudoAdmin(msg *gotgbot.Message) bool {
	tgId := msg.From.Id

	return helper.SudoAdmins[tgId]
}
