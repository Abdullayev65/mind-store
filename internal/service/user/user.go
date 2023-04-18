package user

import (
	"errors"
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/encoder"
	"mindstore/pkg/hash-types"
	"mindstore/pkg/util/timeutil"
)

type Service struct {
	User User
	Auth Auth
}

func New(user User) *Service {
	n := new(Service)

	n.User = user

	return n
}

func (s *Service) UserById(c ctx.Ctx, id hash.Int) (*model.User, error) {
	return s.User.GetById(c, id)
}

func (s *Service) DetailById(c ctx.Ctx, id *hash.Int) (*user.UserDetail, error) {
	obj, err := s.User.DetailById(c, id)
	if err != nil {
		return nil, err
	}

	timeutil.Format(obj.BirthDate, &obj.BirthDateStr)

	return obj, nil
}

func (s *Service) UserUpdate(c ctx.Ctx, input *user.UserUpdate) error {
	var errStr string
	switch {
	case input.Email != nil && (!s.Auth.IsValidEmail(*input.Email)):
		errStr = "email is not valid"
	case input.Username != nil && (len(*input.Username) < 3 || len(*input.Username) > 26):
		errStr = "username length should be between 3 and 26"
	case input.Password != nil && (len(*input.Password) < 3 || len(*input.Password) > 30):
		errStr = "password length should be between 3 and 30"
	}
	if errStr != "" {
		return errors.New(errStr)
	}

	if input.Password != nil {
		password, err := encoder.HashPassword(*input.Password)
		if err != nil {
			return err
		}

		input.Password = &password
	}

	timeutil.Parse(input.BirthDateStr, &input.BirthDate)

	return s.User.Update(c, input)
}
