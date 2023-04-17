package service

import (
	"mindstore/internal/repository"
	"mindstore/internal/service/auth"
	"mindstore/internal/service/user"
)

var Auth = &auth.Service{}
var User = new(user.Service)

func init() {
	*Auth = *auth.New(repository.User)
	*User = *user.New(repository.User)
}
