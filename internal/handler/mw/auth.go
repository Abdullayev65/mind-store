package mw

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	. "mindstore/pkg/response"
	"strings"
)

func (mw *MiddleWere) UserFromToken(required bool) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		user, err := mw.userFromToken(c, authHeader)

		if err == nil {
			c.Set("user", user)
		} else if required {
			FailErr(c, err)
			c.Abort()
			return
		}
	}
}

func (mw *MiddleWere) userFromToken(c ctx.Ctx, authHeader string) (*model.User, error) {
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

	user, err := mw.user.UserById(c, *id)

	return user, err
}
