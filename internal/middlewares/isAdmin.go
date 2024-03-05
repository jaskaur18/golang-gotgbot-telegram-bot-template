package middlewares

import (
	"context"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/constant"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/db"
)

func IsAdmin(s *bot.Server, msg *gotgbot.Message) bool {
	tgID := msg.From.Id

	if IsSudoAdmin(msg) {
		return true
	}

	u, err := s.Queries.GetUserByTelegramID(context.Background(), pgtype.Int8{Int64: tgID, Valid: true})
	if err != nil {
		return false
	}

	if u.UserType == db.UsertypeADMIN {
		return true
	}

	return false
}

func IsSudoAdmin(msg *gotgbot.Message) bool {
	tgID := msg.From.Id

	return constant.GetSudoAdmins()[tgID]
}
