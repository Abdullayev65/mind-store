package user

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/object/dto/user"
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

func (h *Handler) UserGetMe(c *gin.Context) {
	userId := h.authMW.MustGetUserId(c)

	detail, err := h.user.DetailById(c, userId)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, detail)
}

func (h *Handler) UserUpdate(c *gin.Context, input *user.UserUpdate) {
	input.Id = *h.authMW.MustGetUserId(c)

	err := h.user.UserUpdate(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, "DONE")
}
