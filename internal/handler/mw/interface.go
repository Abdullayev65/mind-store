package mw

import (
	"mindstore/pkg/hash-types"
)

type Auth interface {
	UserIdFromToken(tokenStr string) (*hash.Int, error)
}
