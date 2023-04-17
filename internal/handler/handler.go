package handler

import (
	"mindstore/internal/handler/mw"
	"mindstore/internal/handler/v1/auth"
	"mindstore/internal/handler/v1/user"
	"mindstore/internal/service"
)

var (
	MW   = new(mw.MiddleWere)
	Auth = new(auth.Handler)
	User = new(user.Handler)
)

func init() {
	*MW = *mw.New(service.Auth, service.User)
	*Auth = *auth.New(service.Auth)
	*User = *user.New(service.User, MW)
}
