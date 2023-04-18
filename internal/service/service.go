package service

import (
	repo "mindstore/internal/repository"
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
	*Auth = *auth.New(repo.User)
	*User = *user.New(repo.User, Auth)
	*Mind = *mind.New(repo.Mind, nil)
}
