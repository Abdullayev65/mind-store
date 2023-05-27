package bind

import (
	"github.com/gin-gonic/gin"
	"mindstore/pkg/binding"
	"mindstore/pkg/response"
)

const bindingKey = "BiNdInG_KEy"

func Binder[In any](handler func(*gin.Context, *In)) gin.HandlerFunc {
	return func(c *gin.Context) {
		in := new(In)

		b := binding.Default(c.Request.Method, c.ContentType())

		err := c.ShouldBindWith(in, b)
		if err != nil {
			response.FailErr(c, err)
			c.Abort()
			return
		}

		handler(c, in)
	}
}

func BindQuery(c *gin.Context, obj any) error {
	values := c.Request.URL.Query()
	if err := binding.MapForm(obj, values); err != nil {
		return err
	}
	return binding.Validate(obj)
}

//func MustGet[In any](c *gin.Context) *In {
//	in, err := Get[In](c)
//	if err != nil {
//		panic(err)
//	}
//
//	return in
//}
//
//func Get[In any](c *gin.Context) (*In, error) {
//	val, ok := c.Get(bindingKey)
//	if !ok {
//		return nil, errors.New("bind: input data not found")
//	}
//
//	if in, ok := val.(*In); ok {
//		return in, nil
//	} else {
//		return nil, fmt.Errorf("bind: inputdata can't be converted to %T", in)
//	}
//
//}
