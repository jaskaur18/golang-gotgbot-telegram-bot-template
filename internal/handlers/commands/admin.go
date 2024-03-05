package commands

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/db"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/middlewares"
	"github.com/rs/zerolog/log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func HandleAdmin(s *bot.Server, b *gotgbot.Bot, ctx *ext.Context) error {
	if !middlewares.IsSudoAdmin(ctx.EffectiveMessage) {
		return replyToMessage(b, ctx, "You are not allowed to use this command")
	}

	userID, newAdminStatus, err := parseAdminCommandArguments(ctx)
	if err != nil {
		return replyToMessage(b, ctx, err.Error())
	}

	targetUser, err := s.Queries.GetUserByTelegramID(context.Background(), pgtype.Int8{Int64: userID, Valid: true})
	if err != nil {
		logUserNotFound(ctx, userID)
		return replyToMessage(b, ctx, "User not found")
	}

	if err := updateUserAdminStatus(b, ctx, s, &targetUser, newAdminStatus); err != nil {
		log.Error().Err(err).Msg("Error updating user's admin status")
		return replyToMessage(b, ctx, "Error updating user's admin status")
	}

	return nil
}

func parseAdminCommandArguments(ctx *ext.Context) (int64, bool, error) {
	args := ctx.Args()
	requiredArgs := 2
	if len(args) < requiredArgs {
		return 0, false, errors.New("usage: /admin id true/false or /admin @username true/false")
	}

	var userID int64
	var err error
	userIdentifier := args[1]
	if strings.HasPrefix(userIdentifier, "@") && ctx.EffectiveMessage.ReplyToMessage != nil {
		userID = ctx.EffectiveMessage.ReplyToMessage.From.Id
	} else {
		userID, err = strconv.ParseInt(userIdentifier, 10, 64)
		if err != nil {
			return 0, false, errors.New("please provide a valid user identifier (ID or @username)")
		}
	}

	if len(args) < 3 && !strings.HasPrefix(userIdentifier, "@") {
		return 0, false, errors.New("usage: /admin id true/false")
	}

	statusArg := args[len(args)-1]
	newAdminStatus := strings.EqualFold(statusArg, "true")

	return userID, newAdminStatus, nil
}

func logUserNotFound(ctx *ext.Context, userID int64) {
	log.Debug().
		Str("user_id", strconv.FormatInt(userID, 10)).
		Str("adminID", strconv.FormatInt(ctx.EffectiveUser.Id, 10)).
		Msg("User not found")
}

func updateUserAdminStatus(
	b *gotgbot.Bot, ctx *ext.Context, s *bot.Server, targetUser *db.User, newAdminStatus bool) error {
	selectUserType := db.UsertypeUSER
	if newAdminStatus {
		selectUserType = db.UsertypeADMIN
	}

	if targetUser.UserType == selectUserType {
		msg := fmt.Sprintf("User %d is already an %s", targetUser.TelegramID.Int64, adminStatus(newAdminStatus))
		log.Info().Msg(msg)
		return replyToMessage(b, ctx, msg)
	}

	params := db.UpdateUserTypeParams{
		TelegramID: targetUser.TelegramID,
		UserType:   selectUserType,
	}
	if err := s.Queries.UpdateUserType(context.Background(), params); err != nil {
		return err
	}

	msg := fmt.Sprintf("User %d is now an %s", targetUser.TelegramID.Int64, adminStatus(newAdminStatus))
	log.Info().Msg(msg)
	return replyToMessage(b, ctx, msg)
}

func replyToMessage(b *gotgbot.Bot, ctx *ext.Context, text string) error {
	_, err := ctx.EffectiveMessage.Reply(b, text, nil)
	return err
}

func adminStatus(isAdmin bool) string {
	if isAdmin {
		return "admin"
	}
	return "normal user"
}
