package main

import (
	"bot/handlers"
	"bot/helpers"
	"bot/model"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
	"net/http"
)

func init() {
	helpers.InitEnv()
	helpers.InitConstants()
	model.InitDB()
}

func main() {
	bot, err := gotgbot.NewBot(helpers.Env.BotToken, &gotgbot.BotOpts{
		Client: http.Client{},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: gotgbot.DefaultTimeout,
			APIURL:  gotgbot.DefaultAPIURL,
		},
	})
	if err != nil {
		log.Fatal("Error creating bot: ", err)
		return
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			// If an error is returned by a handler, log it and continue going.
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		}),
	})
	handlers.Load(updater.Dispatcher)

	if helpers.Env.PROD {
		helpers.ProdLaunch(bot, updater)
	} else {
		helpers.DevLaunch(bot, updater)
	}

	log.Println("🔥 Bot Is Running 🔥")
	log.Printf("🔗 Bot Username: @%s\n", bot.Username)
	log.Printf("🆔 Admin Ids: %v\n", helpers.Env.AdminIds)
	updater.Idle()
}
