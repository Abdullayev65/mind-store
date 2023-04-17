package auth

import (
	"github.com/gin-gonic/gin"
	dtoauth "mindstore/internal/object/dto/auth"
	dtouser "mindstore/internal/object/dto/user"
	auth_srvc "mindstore/internal/service/auth"
	. "mindstore/pkg/response"
)

type Handler struct {
	auth *auth_srvc.Service
}

func New(auth *auth_srvc.Service) *Handler {
	return &Handler{auth}
}

func (h *Handler) SignUp(c *gin.Context) {
	data := new(dtouser.UserCreate)

	err := c.Bind(data)
	if err != nil {
		FailErr(c, err)
		return
	}

	err = h.auth.SignUp(c, data)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, "user created")
}

func (h *Handler) LogIn(c *gin.Context) {
	data := new(dtoauth.LogIn)

	err := c.Bind(data)
	if err != nil {
		FailErr(c, err)
		return
	}

	outPut, err := h.auth.LogIn(c, data)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, outPut)
}
