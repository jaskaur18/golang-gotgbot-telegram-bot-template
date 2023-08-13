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
	if !helper.SudoAdmins[c.EffectiveUser.Id] {
		_, err := c.EffectiveMessage.Reply(b, "You are not allowed to use this command", nil)
		return err
	}

	args := c.Args()
	if len(args) < 3 {
		_, err := c.EffectiveMessage.Reply(b, "Usage: /admin id true/false", nil)
		return err
	}

	userId := args[0]
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		_, err := c.EffectiveMessage.Reply(b, "Please provide a valid user id", nil)
		return err
	}

	targetUser, err := helper.DB.User.FindFirst(db.User.TelegramID.Equals(userIdInt)).Exec(context.Background())
	if err != nil || targetUser == nil {
		_, err := c.EffectiveMessage.Reply(b, "User not found", nil)
		return err
	}

	newAdminStatusStr := strings.ToLower(args[1])
	if newAdminStatusStr != "true" && newAdminStatusStr != "false" {
		_, err := c.EffectiveMessage.Reply(b, "Invalid admin status, use 'true' or 'false'", nil)
		return err
	}

	newAdminStatus := newAdminStatusStr == "true"
	currentAdminStatus := targetUser.UserType == db.UserTypeADMIN

	if newAdminStatus == currentAdminStatus {
		adminType := "admin"
		if !newAdminStatus {
			adminType = "normal user"
		}
		_, err := c.EffectiveMessage.Reply(b, fmt.Sprintf("User %s is already an %s", userId, adminType), nil)
		return err
	}

	_, err = helper.DB.User.FindMany(db.User.TelegramID.Equals(userIdInt)).Update(
		db.User.UserType.Set(db.UserTypeADMIN),
	).Exec(context.Background())

	if err != nil {
		_, err := c.EffectiveMessage.Reply(b, "Error updating user's admin status", nil)
		return err
	}

	adminType := "admin"
	if !newAdminStatus {
		adminType = "normal user"
	}

	msg := fmt.Sprintf("User %s is now a %s", userId, adminType)
	_, err = c.EffectiveMessage.Reply(b, msg, nil)
	return err
}
