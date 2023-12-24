package menus

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/gotgbot/keyboard"
	"github.com/jaskaur18/moimoiStoreBot/helpers"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

func ServicesKeyboard() *gotgbot.InlineKeyboardMarkup {
	k := new(keyboard.InlineKeyboard)
	prods := []*models.Record{}

	err := helpers.PB.Dao().RecordQuery("product").All(&prods)
	if err != nil {
		if apis.NewNotFoundError(err.Error(), err) != nil {
			return k.Build()
		}
	}

	for _, prod := range prods {
		h := prod.Get("hidden").(bool)
		if h {
			continue
		}

		name := prod.Get("name").(string)
		id := prod.Get("id").(string)
		k.Text(name, fmt.Sprintf("prod:%s", id))
	}

	return k.Build()
}
