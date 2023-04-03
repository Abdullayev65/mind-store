package user

import (
	"context"
	"mindstore/internal/basic"
	"mindstore/internal/model"
)

type Repository interface {
	basic.CRUD[model.User, UserCreate, UserUpdate, Filter]
	GetByUsername(c context.Context, username string) (*model.User, error)
	GetByEmail(c context.Context, email string) (*model.User, error)
}
