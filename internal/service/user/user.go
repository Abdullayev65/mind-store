package user

import (
	"errors"
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/encoder"
	"mindstore/pkg/hash-types"
)

type Service struct {
	User
}

func New(user User) *Service {
	n := new(Service)

	n.User = user

	return n
}

func (s *Service) Create(c ctx.Ctx, data *user.UserCreate) (*model.User, error) {
	var errStr string
	switch {
	case data.Email == nil:
		errStr = "email is required"
	case data.Username == nil:
		errStr = "username is required"
	case data.Password == nil:
		errStr = "password is required"
	case data.FirstName == nil:
		errStr = "first_name is required"
	}
	if errStr != "" {
		return nil, errors.New(errStr)
	}

	password, err := encoder.HashPassword(*data.Password)
	if err != nil {
		return nil, err
	}

	data.Password = &password

	return s.User.Create(c, data)
}

func (s *Service) UserById(c ctx.Ctx, id hash.Int) (*model.User, error) {
	return s.User.GetById(c, id)
}
