package router

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/handler"
	"mindstore/pkg/bind"
)

func Mind(r *gin.RouterGroup) {
	h := handler.Mind
	mw := handler.MW

	r.POST("", mw.UserIdFromToken(true), bind.Binder(h.CreateMind))
	r.PUT("/:mind_id", mw.UserIdFromToken(true), bind.Binder(h.UpdateMind))
	r.DELETE("/:mind_id", mw.UserIdFromToken(true), h.DeleteMind)

	r.GET("/sub-minds/:parent_mind_id", mw.UserIdFromToken(false),
		h.GetMindChildren("parent_mind_id", false))

	r.GET("/with-sub-minds/:mind_id", mw.UserIdFromToken(false),
		h.GetMindChildren("mind_id", true))

	r.POST("/add-file", mw.UserIdFromToken(true), bind.Binder(h.AddFile))
	r.DELETE("/delete-file/:mind_id/:file_id", mw.UserIdFromToken(true), h.DeleteFile)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
