package commands

import (
	"bot/db"
	"bot/helper"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func HandleAdmin(b *gotgbot.Bot, c *ext.Context) error {
	// Check if the user is authorized to use this command
	if !helper.SudoAdmins[c.EffectiveUser.Id] {
		_, _ = c.EffectiveMessage.Reply(b, "You are not allowed to use this command", nil)
		return nil
	}

	// Parse command arguments
	args := c.Args()
	if len(args) < 2 {
		_, _ = c.EffectiveMessage.Reply(b, "Usage: /admin id true/false", nil)
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
			_, _ = c.EffectiveMessage.Reply(b, "Usage: /admin id true/false", nil)
			return nil
		}
	} else if strings.HasPrefix(userIdentifier, "@") {
		// Handle the case when the identifier starts with @
		// If replying to a forwarded message, get user ID from the forwarded message
		if c.EffectiveMessage.ReplyToMessage != nil {
			userIdInt = c.EffectiveMessage.ReplyToMessage.From.Id
			if len(args) >= 2 {
				newAdminStatusStr = strings.ToLower(args[1])
			} else {
				_, _ = c.EffectiveMessage.Reply(b, "Usage: /admin @username true/false", nil)
				return nil
			}
		} else {
			_, _ = c.EffectiveMessage.Reply(b, "Please reply to a forwarded message to identify the user", nil)
			return nil
		}
	} else {
		_, _ = c.EffectiveMessage.Reply(b, "Please provide a valid user identifier (ID or @username)", nil)
		return nil
	}

	// Find the target user
	targetUser, err := helper.DB.User.FindFirst(db.User.TelegramID.Equals(db.BigInt(userIdInt))).Exec(context.Background())
	if err != nil || targetUser == nil {
		_, _ = c.EffectiveMessage.Reply(b, "User not found", nil)
		return nil
	}

	// Parse the new admin status
	if newAdminStatusStr != "true" && newAdminStatusStr != "false" {
		_, _ = c.EffectiveMessage.Reply(b, "Invalid admin status, use 'true' or 'false'", nil)
		return nil
	}
	newAdminStatus := newAdminStatusStr == "true"
	currentAdminStatus := targetUser.UserType == db.UserTypeAdmin

	// Check if the new status is the same as the current status
	if newAdminStatus == currentAdminStatus {
		adminType := "admin"
		if !newAdminStatus {
			adminType = "normal user"
		}
		_, _ = c.EffectiveMessage.Reply(b, fmt.Sprintf("User %s is already an %s", userIdentifier, adminType), nil)
		return nil
	}

	// Update the user's admin status
	_, err = helper.DB.User.FindMany(db.User.TelegramID.Equals(db.BigInt(userIdInt))).Update(
		db.User.UserType.Set(db.UserTypeAdmin),
	).Exec(context.Background())

	if err != nil {
		_, _ = c.EffectiveMessage.Reply(b, "Error updating user's admin status", nil)
		return nil
	}

	// Construct and send success message
	adminType := "admin"
	if !newAdminStatus {
		adminType = "normal user"
	}
	msg := fmt.Sprintf("User %s is now a %s", userIdentifier, adminType)
	_, _ = c.EffectiveMessage.Reply(b, msg, nil)
	return nil
}
