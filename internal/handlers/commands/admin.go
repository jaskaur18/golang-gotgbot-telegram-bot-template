package commands

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/cmd/bot"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/db"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/helper"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func HandleAdmin(s *bot.Server, b *gotgbot.Bot, ctx *ext.Context) error {
	// Check if the user is authorized to use this command
	if !helper.SudoAdmins[ctx.EffectiveUser.Id] {
		_, _ = ctx.EffectiveMessage.Reply(b, "You are not allowed to use this command", nil)
		return nil
	}

	// Parse command arguments
	args := ctx.Args()
	if len(args) < 2 {
		_, _ = ctx.EffectiveMessage.Reply(b, "Usage: /admin id true/false", nil)
		return nil
	}

	// Parse the target user identifier
	userIdentifier := args[1]
	var userIdInt int64
	newAdminStatusStr := ""

	// Check if userIdentifier is an integer (user ID)
	if id, err := strconv.ParseInt(userIdentifier, 10, 64); err == nil {
		userIdInt = id
		if len(args) >= 3 {
			newAdminStatusStr = strings.ToLower(args[2])
		} else {
			_, _ = ctx.EffectiveMessage.Reply(b, "Usage: /admin id true/false", nil)
			return nil
		}
	} else if strings.HasPrefix(userIdentifier, "@") {
		// Handle the case when the identifier starts with @
		// If replying to a forwarded message, get user ID from the forwarded message
		if ctx.EffectiveMessage.ReplyToMessage != nil {
			userIdInt = ctx.EffectiveMessage.ReplyToMessage.From.Id
			if len(args) >= 2 {
				newAdminStatusStr = strings.ToLower(args[1])
			} else {
				_, _ = ctx.EffectiveMessage.Reply(b, "Usage: /admin @username true/false", nil)
				return nil
			}
		} else {
			_, _ = ctx.EffectiveMessage.Reply(b, "Please reply to a forwarded message to identify the user", nil)
			return nil
		}
	} else {
		_, _ = ctx.EffectiveMessage.Reply(b, "Please provide a valid user identifier (ID or @username)", nil)
		return nil
	}

	targetUser, err := s.Queries.GetUserByTelegramID(context.Background(), pgtype.Int8{Int64: userIdInt, Valid: true})
	if err != nil {
		log.Debug().Str("user_id", strconv.FormatInt(userIdInt, 10)).
			Str("adminID", strconv.FormatInt(ctx.EffectiveUser.Id, 10)).Msg("User not found")

		_, _ = ctx.EffectiveMessage.Reply(b, "User not found", nil)
		return nil
	}

	// Parse the new admin status
	if newAdminStatusStr != "true" && newAdminStatusStr != "false" {
		_, _ = ctx.EffectiveMessage.Reply(b, "Invalid admin status, use 'true' or 'false'", nil)
		return nil
	}

	newAdminStatus := newAdminStatusStr == "true"
	currentAdminStatus := targetUser.UserType == db.UsertypeADMIN

	// Check if the new status is the same as the current status
	if newAdminStatus == currentAdminStatus {
		adminType := "admin"
		if !newAdminStatus {
			adminType = "normal user"
		}
		log.Debug().Str("user_id", strconv.FormatInt(userIdInt, 10)).
			Str("adminID", strconv.FormatInt(ctx.EffectiveUser.Id, 10)).
			Str("adminType", adminType).Msg("User is already an admin")

		_, _ = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("User %s is already an %s", userIdentifier, adminType), nil)
		return nil
	}

	targetUser.UserType = db.UsertypeADMIN

	// Update the user's admin status
	params := db.UpdateUserTypeParams{
		TelegramID: pgtype.Int8{Int64: userIdInt, Valid: true},
		UserType:   db.UsertypeADMIN,
	}

	if err := s.Queries.UpdateUserType(context.Background(), params); err != nil {
		log.Debug().Str("user_id", strconv.FormatInt(userIdInt, 10)).
			Str("adminID", strconv.FormatInt(ctx.EffectiveUser.Id, 10)).
			Err(err).Msg("Error updating user's admin status")

		_, _ = ctx.EffectiveMessage.Reply(b, "Error updating user's admin status", nil)
		return nil
	}

	// Construct and send a success message
	adminType := "admin"
	if !newAdminStatus {
		adminType = "normal user"
	}

	log.Info().Int64("user_id", userIdInt).
		Int64("adminID", ctx.EffectiveUser.Id).
		Str("adminType", adminType).Msg("User is now an admin")

	msg := fmt.Sprintf("User %s is now a %s", userIdentifier, adminType)
	_, _ = ctx.EffectiveMessage.Reply(b, msg, nil)
	return nil
}
