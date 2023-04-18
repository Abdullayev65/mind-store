package router

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/handler"
	"mindstore/pkg/bind"
)

func Mind(r *gin.RouterGroup) {
	h := handler.Mind
	mw := handler.MW

	r.POST("/", mw.UserIdFromToken(true), bind.Binder(h.CreateMind))
	//r.PUT("me", mw.UserIdFromToken(true), bind.Binder(h.UserUpdate))
	//r.DELETE("me", mw.UserIdFromToken(true), h.UserDelete)

}
