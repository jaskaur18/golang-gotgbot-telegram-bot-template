package utils

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/rs/zerolog/log"
)

func ProdLaunch(_ *gotgbot.Bot, _ *ext.Updater) error {
	// Start the webhook server. We start the server before we set the webhook itself, so that when telegram starts
	// sending updates, the server is already ready.
	// webhookOpts := ext.WebhookOpts{
	//	ListenAddr: "localhost:8080", // This example assumes you're in a dev environment running ngrok on 8080.
	//}
	// err := updater.StartWebhook(b, Env.BotToken, webhookOpts)
	// if err != nil {
	//	panic("failed to start webhook: " + err.Error())
	//}
	//
	// err = updater.SetAllBotWebhooks(Env.WebhookUrl, &gotgbot.SetWebhookOpts{
	//	MaxConnections:     100,
	//	DropPendingUpdates: false,
	// })
	// if err != nil {
	//	panic("failed to set webhook: " + err.Error())
	//}
	// log.Printf("🔗 Webhook URL: %s\n", Env.WebhookUrl)
	return nil
}

func DevLaunch(b *gotgbot.Bot, updater *ext.Updater) error {
	err := updater.StartPolling(
		b, &ext.PollingOpts{
			DropPendingUpdates:    true,
			EnableWebhookDeletion: true,
		},
	)
	if err != nil {
		log.Err(err).Msg("Failed to start polling")
		return err
	}

	return nil
}
