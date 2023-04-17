package router

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/handler"
)

func User(r *gin.RouterGroup) {
	h := handler.User
	mw := handler.MW

	r.GET("me", mw.UserIdFromToken(true), h.Me)

}
