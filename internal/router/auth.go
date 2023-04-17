package router

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/handler"
)

func Auth(r *gin.RouterGroup) {
	h := handler.Auth

	r.POST("sign-up", h.SignUp)
	r.POST("log-in", h.LogIn)
}
