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
	r.PUT("/", mw.UserIdFromToken(true), bind.Binder(h.UpdateMind))
	//r.GET("/children/:parent_mind_id", mw.UserIdFromToken(true), h.GetMindChildren("parent_mind_id", false))
	//r.GET("/with-children/:parent_mind_id", mw.UserIdFromToken(true), h.GetMindChildren("parent_mind_id", true))

	//r.DELETE("me", mw.UserIdFromToken(true), h.UserDelete)

}
