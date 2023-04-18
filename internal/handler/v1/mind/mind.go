package mind

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/object/dto/mind"
	mind_srvc "mindstore/internal/service/mind"
	. "mindstore/pkg/response"
)

type Handler struct {
	mind   *mind_srvc.Service
	authMW AuthMW
}

func New(user *mind_srvc.Service, authMW AuthMW) *Handler {
	return &Handler{user, authMW}
}

func (h *Handler) CreateMind(c *gin.Context, input *mind.Create) {
	input.CreatedBy = h.authMW.MustGetUserId(c)

	id, err := h.mind.CreateMind(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, &id)
}
