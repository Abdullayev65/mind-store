package auth

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/object/dto/auth"
	"mindstore/internal/object/dto/user"
	auth_srvc "mindstore/internal/service/auth"
	. "mindstore/pkg/response"
)

type Handler struct {
	auth *auth_srvc.Service
}

func New(auth *auth_srvc.Service) *Handler {
	return &Handler{auth}
}

func (h *Handler) SignUp(c *gin.Context, input *user.UserCreate) {
	password := input.Password

	err := h.auth.SignUp(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	login := &auth.LogIn{
		Identifier: input.Username,
		Password:   password,
	}

	outPut, err := h.auth.LogIn(c, login)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, outPut)
}

func (h *Handler) LogIn(c *gin.Context, input *auth.LogIn) {
	outPut, err := h.auth.LogIn(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, outPut)
}

func (h *Handler) Available(c *gin.Context, input *auth.Available) {
	outPut, err := h.auth.Available(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, outPut)
}
