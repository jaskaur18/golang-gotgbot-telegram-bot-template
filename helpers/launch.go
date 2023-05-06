package helpers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
)

func ProdLaunch(b *gotgbot.Bot, updater *ext.Updater) {
	// Start the webhook server. We start the server before we set the webhook itself, so that when telegram starts
	// sending updates, the server is already ready.
	webhookOpts := ext.WebhookOpts{
		ListenAddr: "localhost:8080", // This example assumes you're in a dev environment running ngrok on 8080.
	}
	// We use the token as the urlPath for the webhook, as using a secret ensures that strangers aren't crafting fake updates.
	err := updater.StartWebhook(b, Env.BotToken, webhookOpts)
	if err != nil {
		panic("failed to start webhook: " + err.Error())
	}

	err = updater.SetAllBotWebhooks(Env.WebhookUrl, &gotgbot.SetWebhookOpts{
		MaxConnections:     100,
		DropPendingUpdates: false,
	})
	if err != nil {
		panic("failed to set webhook: " + err.Error())
	}
	log.Printf("🔗 Webhook URL: %s\n", Env.WebhookUrl)

}

func DevLaunch(b *gotgbot.Bot, updater *ext.Updater) {
	_, err := b.DeleteWebhook(&gotgbot.DeleteWebhookOpts{
		DropPendingUpdates: true,
	})
	if err != nil {
		log.Panic("Error deleting webhook: ", err)
	}
	err = updater.StartPolling(
		b, &ext.PollingOpts{
			// DropPendingUpdates: true,
		},
	)
	if err != nil {
		log.Fatal("Error starting updater: ", err)
		return
	}
}
