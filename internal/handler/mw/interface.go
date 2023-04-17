package mw

import (
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
)

type Auth interface {
	UserIdFromToken(tokenStr string) (*hash.Int, error)
}

type User interface {
	UserById(c ctx.Ctx, id hash.Int) (*model.User, error)
}
