package response

import (
	"github.com/gin-gonic/gin"
	"mindstore/pkg/render"
)

func Success(c *gin.Context, res any) {
	json200(c, map[string]interface{}{
		"res":    res,
		"status": true,
	})
}

func Fail(c *gin.Context, msg string) {
	json200(c, map[string]interface{}{
		"status": false,
		"msg":    msg,
	})
}

func FailErr(c *gin.Context, err error) {
	Fail(c, err.Error())
}

func json200(c *gin.Context, output any) {
	c.Render(200, render.JSON{Data: output})
}
