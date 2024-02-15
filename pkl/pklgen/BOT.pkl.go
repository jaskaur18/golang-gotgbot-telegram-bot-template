// Code generated from Pkl module `botConfig.pkl`. DO NOT EDIT.
package pklgen

type BOT struct {
	Token string `pkl:"Token"`

	SudoAdmins []int32 `pkl:"SudoAdmins"`

	WebhookUrl *string `pkl:"WebhookUrl"`

	WebhookSecret *string `pkl:"WebhookSecret"`
}
