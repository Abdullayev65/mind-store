package user

import (
	"github.com/gin-gonic/gin"
	user_srvc "mindstore/internal/service/user"
	. "mindstore/pkg/response"
)

type Handler struct {
	user   *user_srvc.Service
	authMW AuthMW
}

func New(user *user_srvc.Service, authMW AuthMW) *Handler {
	return &Handler{user, authMW}
}

func (h *Handler) Me(c *gin.Context) {
	userId := h.authMW.MustGetUserId(c)

	detail, err := h.user.DetailById(c, userId)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, detail)
}
