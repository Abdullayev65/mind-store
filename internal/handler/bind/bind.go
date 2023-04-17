package bind

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mindstore/pkg/binding"
	"mindstore/pkg/response"
)

const bindingKey = "BiNdInG_KEy"

func Binder[In any](c *gin.Context) {
	obj := new(In)

	b := binding.Default(c.Request.Method, c.ContentType())

	err := c.ShouldBindWith(obj, b)
	if err != nil {
		response.FailErr(c, err)
		c.Abort()
		return
	}

	c.Set(bindingKey, obj)
}

func MustGet[In any](c *gin.Context) *In {
	in, err := Get[In](c)
	if err != nil {
		panic(err)
	}

	return in
}

func Get[In any](c *gin.Context) (*In, error) {
	val, ok := c.Get(bindingKey)
	if !ok {
		return nil, errors.New("bind: input data not found")
	}

	if in, ok := val.(*In); ok {
		return in, nil
	} else {
		return nil, fmt.Errorf("bind: inputdata can't be converted to %T", in)
	}

}
