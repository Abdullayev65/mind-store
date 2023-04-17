package mw

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mindstore/pkg/hash-types"
	. "mindstore/pkg/response"
	"strings"
)

func (mw *MiddleWere) UserIdFromToken(required bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			if required {
				Fail(c, "Authorization header is required")
				c.Abort()
			}
			return
		}

		id, err := mw.userIdFromToken(authHeader)

		if err == nil {
			c.Set(userIdKey, id)
		} else if required {
			FailErr(c, err)
			c.Abort()
			return
		}
	}
}

func (mw *MiddleWere) userIdFromToken(authHeader string) (*hash.Int, error) {
	fields := strings.Fields(authHeader)
	switch {
	case len(fields) == 0:
		return nil, errors.New("Authorization token not given")
	case len(fields) > 2:
		return nil, errors.New("invalid Authorization token")
	}

	tokenStr := fields[len(fields)-1]
	id, err := mw.auth.UserIdFromToken(tokenStr)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (mw *MiddleWere) GetUserId(c *gin.Context) (id *hash.Int, ok bool) {
	val, exists := c.Get(userIdKey)
	if !exists {
		return nil, false
	}

	id, ok = val.(*hash.Int)
	return id, ok
}

func (mw *MiddleWere) MustGetUserId(c *gin.Context) *hash.Int {
	id, ok := mw.GetUserId(c)
	if !ok {
		panic("auth: userId not fount")
	}

	return id
}
