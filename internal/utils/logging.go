package utils

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func SetupLogger(level zerolog.Level, prettyPrint bool) {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack //nolint:reassign
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.SetGlobalLevel(level)

	if prettyPrint {
		log.Logger = log.Output(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.TimeFormat = "15:04:05"
		}))
	}
}
