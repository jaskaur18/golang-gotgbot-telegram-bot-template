package main

import (
	"bot/handlers"
	"bot/helpers"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
)

func init() {
	helpers.InitEnv()
	helpers.InitConstants()
}

func main() {
	bot, err := gotgbot.NewBot(helpers.Env.BotToken, nil)
	if err != nil {
		log.Fatal("Error creating bot: ", err)
		return
	}
	updater := ext.NewUpdater(nil)
	handlers.Load(updater.Dispatcher)
	err = updater.StartPolling(
		bot, &ext.PollingOpts{
			// DropPendingUpdates: true,
		},
	)
	if err != nil {
		log.Fatal("Error starting updater: ", err)
		return
	}

	log.Println("🔥 Bot Is Running 🔥")
	log.Printf("🔗 Bot Username: @%s\n", bot.Username)
	log.Printf("🆔 Admin Ids: %v\n", helpers.Env.AdminIds)
	updater.Idle()
}
