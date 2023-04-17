package handler

import (
	"mindstore/internal/handler/mw"
	"mindstore/internal/handler/v1/auth"
	"mindstore/internal/service"
)

var (
	Auth = new(auth.Handler)
	MW   = new(mw.MiddleWere)
)

func init() {
	*Auth = *auth.New(service.Auth)
	*MW = *mw.New(service.Auth, service.User)
}
