package helpers

import "github.com/PaulSonOfLars/gotgbot/v2/ext"

var CustomContext = ext.Context{
	Data: map[string]interface{}{
		"isAdmin": false,
	},
}
