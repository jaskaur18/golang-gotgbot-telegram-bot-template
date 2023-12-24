package menus

import "github.com/gotgbot/keyboard"

var HomeKeyboard = new(keyboard.Keyboard).
	Text("🛒Services").
	Text("👥 Referrals").
	Row().
	Text("🆘 Admin Support").
	Text("Store Policy").Build()

func init() {
	HomeKeyboard.ResizeKeyboard = true
	HomeKeyboard.InputFieldPlaceholder = "Select an option"
}
