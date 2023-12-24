package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/jaskaur18/moimoiStoreBot/bots/storeBot/handlers/filter"
)

type TextFilter struct {
	Text     string
	LevelReq AccessLevel
	Handler  func(b *gotgbot.Bot, ctx *ext.Context) error
}

var TextFilters = []TextFilter{
	{
		Text:     "🛒Services",
		LevelReq: User,
		Handler:  filter.HandleServices,
	},
	{
		Text:     "👥 Referrals",
		LevelReq: User,
		Handler:  filter.HandleReferrals,
	},
	{
		Text:     "🆘 Admin Support",
		LevelReq: User,
		Handler:  filter.HandleSupport,
	},
	{
		Text:     "Store Policy",
		LevelReq: User,
		Handler:  filter.HandleStorePolicy,
	},
}
