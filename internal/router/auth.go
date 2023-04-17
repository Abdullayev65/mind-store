package router

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/handler"
	"mindstore/internal/object/dto/auth"
	"mindstore/internal/object/dto/user"
	"mindstore/pkg/bind"
)

func Auth(r *gin.RouterGroup) {
	h := handler.Auth

	r.POST("sign-up", bind.Binder[user.UserCreate], h.SignUp)
	r.POST("log-in", bind.Binder[auth.LogIn], h.LogIn)
}
