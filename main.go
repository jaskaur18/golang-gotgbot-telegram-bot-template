package main

import (
	"bot/handlers"
	"bot/helper"
	"bot/middlewares"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
	"log/slog"
	"os"
)

func init() {
	helper.InitEnv()
	helper.InitConstants()
	helper.NewDatabase()
	helper.InitRedis()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {

	bot, err := gotgbot.NewBot(helper.Env.BotToken, &gotgbot.BotOpts{
		BotClient: middlewares.NewI18nClient(),
	})
	if err != nil {
		log.Fatal("Error creating bot: ", err)
		return
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			// If a handler returns an error, log it and continue going.
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		}),
	})

	handlers.LoadHandlers(updater.Dispatcher)

	if helper.Env.PROD {
		helper.ProdLaunch(bot, updater)
	} else {
		helper.DevLaunch(bot, updater)
	}

	log.Println("🔥 Bot Is Running 🔥")
	log.Printf("🔗 Bot Username: @%s\n", bot.Username)
	log.Printf("🆔 SudoAdmins: %v\n", helper.Env.SudoAdmins)
	updater.Idle()
}
