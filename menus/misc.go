package menus

import "github.com/gotgbot/keyboard"

var GenderKeyboard = new(keyboard.Keyboard).
	Text("Male").
	Row().
	Text("Female").
	Row().
	Text("Prefer Not To Share").Build()

func init() {
	GenderKeyboard.ResizeKeyboard = true
	GenderKeyboard.OneTimeKeyboard = true
	GenderKeyboard.InputFieldPlaceholder = "Choose Your Gender"
}
