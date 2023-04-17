package auth

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/object/dto/auth"
	"mindstore/internal/object/dto/user"
	auth_srvc "mindstore/internal/service/auth"
	"mindstore/pkg/bind"
	. "mindstore/pkg/response"
)

type Handler struct {
	auth *auth_srvc.Service
}

func New(auth *auth_srvc.Service) *Handler {
	return &Handler{auth}
}

func (h *Handler) SignUp(c *gin.Context) {
	input := bind.MustGet[user.UserCreate](c)

	err := h.auth.SignUp(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, "user created")
}

func (h *Handler) LogIn(c *gin.Context) {
	input := bind.MustGet[auth.LogIn](c)

	outPut, err := h.auth.LogIn(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, outPut)
}
