// Code generated from Pkl module `botConfig.pkl`. DO NOT EDIT.
package levels

import (
	"encoding"
	"fmt"
)

type Levels string

const (
	Debug    Levels = "debug"
	Info     Levels = "info"
	Warn     Levels = "warn"
	Error    Levels = "error"
	Fatal    Levels = "fatal"
	Panic    Levels = "panic"
	Trace    Levels = "trace"
	Disabled Levels = "disabled"
	Nolevel  Levels = "nolevel"
)

// String returns the string representation of Levels
func (rcv Levels) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(Levels)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Levels.
func (rcv *Levels) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "debug":
		*rcv = Debug
	case "info":
		*rcv = Info
	case "warn":
		*rcv = Warn
	case "error":
		*rcv = Error
	case "fatal":
		*rcv = Fatal
	case "panic":
		*rcv = Panic
	case "trace":
		*rcv = Trace
	case "disabled":
		*rcv = Disabled
	case "nolevel":
		*rcv = Nolevel
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid Levels`, str)
	}
	return nil
}
