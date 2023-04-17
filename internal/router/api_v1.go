package router

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/handler"
)

func InitApi() *gin.Engine {
	r := gin.Default()

	r.Use(handler.MW.ErrorHandler)

	v1 := r.Group("/api/v1")

	Auth(v1.Group("/auth"))

	return r
}
