// Code generated from Pkl module `botConfig.pkl`. DO NOT EDIT.
package pklgen

import "github.com/apple/pkl-go/pkl"

type Database struct {
	PSQLDB string `pkl:"PSQLDB"`

	PSQLHOST string `pkl:"PSQLHOST"`

	PSQLPORT int32 `pkl:"PSQLPORT"`

	PSQLUSER string `pkl:"PSQLUSER"`

	PSQLPASS string `pkl:"PSQLPASS"`

	AdditionalParams map[string]string `pkl:"AdditionalParams"`

	DBMaxOpenConns int32 `pkl:"DBMaxOpenConns"`

	MaxIdleConns int32 `pkl:"MaxIdleConns"`

	MinIdleConns int32 `pkl:"MinIdleConns"`

	ConnectionMaxLifetime *pkl.Duration `pkl:"ConnectionMaxLifetime"`
}
