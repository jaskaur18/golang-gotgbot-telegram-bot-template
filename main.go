package bot

import (
<<<<<<< HEAD
=======
	"bot/handlers"
	"bot/helper"
>>>>>>> parent of dc24b0d (Update i18n implementation, libraries and installation script)
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/helper"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/middlewares"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func init() {
	helper.InitEnv()
	helper.InitConstants()
	helper.InitRedis()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

var BOT *gotgbot.Bot

func InitStoreBot() {

	bot, err := gotgbot.NewBot(helper.Env.BotToken, &gotgbot.BotOpts{
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
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
<<<<<<< HEAD

	// Create updater and dispatcher.
	updater := ext.NewUpdater(dispatcher, nil)

	handlers.LoadHandlers(dispatcher)
=======
	handlers.LoadHandlers(updater.Dispatcher)
>>>>>>> parent of dc24b0d (Update i18n implementation, libraries and installation script)

	if helper.Env.PROD {
		helper.ProdLaunch(bot, updater)
	} else {
		helper.DevLaunch(bot, updater)
	}

	BOT = bot

	log.Println("🔥 Bot Is Running 🔥")
	log.Printf("🔗 Bot Username: @%s\n", bot.Username)
	log.Printf("🆔 SudoAdmins: %v\n", helper.Env.SudoAdmins)
	updater.Idle()
}
