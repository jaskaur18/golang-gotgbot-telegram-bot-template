// Code generated from Pkl module `botConfig.pkl`. DO NOT EDIT.
package pklgen

import "github.com/jaskaur18/golang-gotgbot-telegram-bot-template/pkl/pklgen/levels"

type LOGGER struct {
	Level levels.Levels `pkl:"Level"`

	RequestLevel levels.Levels `pkl:"RequestLevel"`

	LogRequestBody bool `pkl:"LogRequestBody"`

	LogRequestHeaders bool `pkl:"LogRequestHeaders"`

	LogRequestQuery bool `pkl:"LogRequestQuery"`

	LogResponseBody bool `pkl:"LogResponseBody"`

	LogResponseHeaders bool `pkl:"LogResponseHeaders"`

	LogCaller bool `pkl:"LogCaller"`

	PrettyPrintConsole bool `pkl:"PrettyPrintConsole"`
}
