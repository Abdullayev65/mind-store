package user

import (
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
)

type User interface {
	GetByUsername(c ctx.Ctx, username string) (*model.User, error)
	GetByEmail(c ctx.Ctx, email string) (*model.User, error)
	Create(ctx.Ctx, *user.UserCreate) (*model.User, error)
	GetById(ctx.Ctx, hash.Int) (*model.User, error)
	DetailById(c ctx.Ctx, id *hash.Int) (*user.UserDetail, error)
	Update(c ctx.Ctx, input *user.UserUpdate) error
	Delete(c ctx.Ctx, userId hash.Int, deletedBy hash.Int) error
}

type Auth interface {
	IsValidEmail(email string) bool
}
