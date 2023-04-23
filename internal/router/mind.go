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
	r.PUT("/:mind_id", mw.UserIdFromToken(true), bind.Binder(h.UpdateMind))

	r.GET("/children/:parent_mind_id", mw.UserIdFromToken(false),
		h.GetMindChildren("parent_mind_id", false))

	r.GET("/with-children/:parent_mind_id", mw.UserIdFromToken(false),
		h.GetMindChildren("parent_mind_id", true))

	r.POST("/add-file", mw.UserIdFromToken(true), bind.Binder(h.AddFile))
	r.DELETE("/delete-file", mw.UserIdFromToken(true), bind.Binder(h.DeleteFile))

}
