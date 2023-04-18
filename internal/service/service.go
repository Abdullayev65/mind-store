package service

import (
	"mindstore/internal/repository"
	"mindstore/internal/service/auth"
	"mindstore/internal/service/mind"
	"mindstore/internal/service/user"
)

var (
	Auth = new(auth.Service)
	User = new(user.Service)
	Mind = new(mind.Service)
)

func init() {
	*Auth = *auth.New(repository.User)
	*User = *user.New(repository.User, Auth)
	*Mind = *mind.New(repository.Mind)
}
