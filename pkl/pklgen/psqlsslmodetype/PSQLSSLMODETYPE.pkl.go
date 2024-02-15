// Code generated from Pkl module `BotConfig.pkl`. DO NOT EDIT.
package psqlsslmodetype

import (
	"encoding"
	"fmt"
)

type PSQLSSLMODETYPE string

const (
	Disable    PSQLSSLMODETYPE = "disable"
	Require    PSQLSSLMODETYPE = "require"
	VerifyCa   PSQLSSLMODETYPE = "verify-ca"
	VerifyFull PSQLSSLMODETYPE = "verify-full"
)

// String returns the string representation of PSQLSSLMODETYPE
func (rcv PSQLSSLMODETYPE) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(PSQLSSLMODETYPE)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for PSQLSSLMODETYPE.
func (rcv *PSQLSSLMODETYPE) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "disable":
		*rcv = Disable
	case "require":
		*rcv = Require
	case "verify-ca":
		*rcv = VerifyCa
	case "verify-full":
		*rcv = VerifyFull
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid PSQLSSLMODETYPE`, str)
	}
	return nil
}
