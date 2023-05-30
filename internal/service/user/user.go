package user

import (
	"errors"
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
	"mindstore/internal/tools/url_tool"
	"mindstore/pkg/config"
	"mindstore/pkg/ctx"
	"mindstore/pkg/encoder"
	"mindstore/pkg/hash-types"
	"mindstore/pkg/util/timeutil"
	"path"
)

type Service struct {
	user       User
	auth       Auth
	systemFile SysFile
	file       File
}

func New(user User, auth Auth, systemFile SysFile, file File) *Service {
	n := new(Service)

	n.user = user
	n.auth = auth
	n.systemFile = systemFile
	n.file = file

	return n
}

func (s *Service) UserById(c ctx.Ctx, id hash.Int) (*model.User, error) {
	return s.user.GetById(c, id)
}

func (s *Service) DetailById(c ctx.Ctx, id *hash.Int) (*user.UserDetail, error) {
	obj, err := s.user.DetailById(c, id)
	if err != nil {
		return nil, err
	}

	timeutil.Format(obj.BirthDate, &obj.BirthDateStr)
	if obj.AvatarId != nil {
		obj.AvatarUrl = url_tool.AvatarUrlWithHash(obj.Id)
	}

	return obj, nil
}

func (s *Service) UserUpdate(c ctx.Ctx, input *user.UserUpdate) error {
	var errStr string
	switch {
	case input.Email != nil && (!s.auth.IsValidEmail(*input.Email)):
		errStr = "email is not valid"
	case input.Username != nil && s.auth.IsValidUsername(*input.Username) != nil:
		errStr = s.auth.IsValidUsername(*input.Username).Error()

	case input.Password != nil && (len(*input.Password) < 1 || len(*input.Password) > 30):
		errStr = "password length should be between 1 and 30"
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
	if input.Avatar != nil {
		fileData, err := s.systemFile.UploadFile(input.Avatar, "avatar")
		if err != nil {
			return err
		}
		fileData.CreatedBy = input.Id
		err = s.file.Create(c, fileData)
		if err != nil {
			return err
		}
		input.AvatarId = &fileData.Id
	}

	timeutil.Parse(input.BirthDateStr, &input.BirthDate)

	return s.user.Update(c, input)
}

func (s *Service) Delete(c ctx.Ctx, input *user.UserDelete) error {
	if input.Password == nil {
		return errors.New("password not given")
	}
	uzer, err := s.user.GetById(c, input.UserId)
	if err != nil {
		return err
	}

	if !encoder.ComparePassword(uzer.Password, *input.Password) {
		return errors.New("password incorrect")
	}

	return s.user.Delete(c, input.UserId, *input.DeletedBy)
}

func (s *Service) UserSearch(c ctx.Ctx, input *user.UserSearch) ([]*user.UserDetail, int, error) {
	switch input.OrderBy {
	case "username":
	case "email":
	case "first_name":
	case "middle_name":
	case "last_name":
	case "id":
	default:
		input.OrderBy = ""
	}
	if input.OrderBy == "" {
		input.OrderBy = "id"
	}
	if input.DescendingOrder {
		input.OrderBy += " DESC"
	}

	details, count, err := s.user.UserSearch(c, input)
	if err != nil {
		return nil, count, err
	}
	for _, d := range details {
		timeutil.Format(d.BirthDate, &d.BirthDateStr)
		if d.AvatarId != nil {
			d.AvatarUrl = url_tool.AvatarUrlWithHash(d.Id)
		}
	}

	return details, count, nil
}

func (s *Service) UserByUsername(c ctx.Ctx, username string) (*user.UserDetail, error) {
	userModel, err := s.user.GetByUsername(c, username)
	if err != nil {
		return nil, err
	}

	userDetail := new(user.UserDetail)

	userDetail.Id = userModel.Id
	userDetail.Username = &userModel.Username
	userDetail.Email = userModel.Email
	userDetail.MindId = userModel.MindId
	userDetail.FirstName = userModel.FirstName
	userDetail.MiddleName = userModel.MiddleName
	userDetail.LastName = userModel.LastName
	timeutil.Parse(userDetail.BirthDateStr, &userModel.BirthDate)

	if userModel.AvatarId != nil {
		url := config.GetFilesUrlWith(path.Join("avatar", userModel.Id.HashToStr()))
		userDetail.AvatarUrl = &url
	}

	return userDetail, nil
}
