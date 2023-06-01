package mind

import (
	"github.com/gin-gonic/gin"
	"mindstore/internal/object/dto/file"
	"mindstore/internal/object/dto/mind"
	file_srvc "mindstore/internal/service/file"
	mind_srvc "mindstore/internal/service/mind"
	. "mindstore/pkg/response"
)

type Handler struct {
	authMW AuthMW
	mind   *mind_srvc.Service
	file   *file_srvc.Service
}

func New(authMW AuthMW, user *mind_srvc.Service, file *file_srvc.Service) *Handler {
	return &Handler{authMW, user, file}
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

func (h *Handler) UpdateMind(c *gin.Context, input *mind.Update) {
	input.CreatedBy = h.authMW.MustGetUserId(c)
	err := input.Id.UnhashStr(c.Param("mind_id"))
	if err != nil {
		FailErr(c, err)
		return
	}

	err = h.mind.UpdateMind(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, "OK")
}

func (h *Handler) GetMindChildren(paramParentMindId string, getOwn bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := new(mind.ChildrenFilter)

		if err := filter.MindId.UnhashStr(c.Param(paramParentMindId)); err != nil {
			FailErr(c, err)
			return
		}
		if userId, ok := h.authMW.GetUserId(c); ok {
			filter.UserId = userId
		}

		if getOwn {
			dto, err := h.mind.WithChildrenById(c, filter)
			if err != nil {
				FailErr(c, err)
				return
			}
			Success(c, dto)
		} else {
			list, err := h.mind.ChildrenById(c, filter)
			if err != nil {
				FailErr(c, err)
				return
			}
			Success(c, list)
		}
	}
}

func (h *Handler) DeleteMind(c *gin.Context) {
	input := new(mind.Delete)
	input.DeleteBy = *h.authMW.MustGetUserId(c)
	input.CreatedBy = input.DeleteBy
	err := input.Id.UnhashStr(c.Param("mind_id"))
	if err != nil {
		FailErr(c, err)
		return
	}

	err = h.mind.Delete(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, "DELETED")
}

func (h *Handler) AddFile(c *gin.Context, input *file.CreateWithMind) {
	if input.MindId == nil {
		Fail(c, "mind_id not found")
		return
	}

	input.CreatedBy = h.authMW.MustGetUserId(c)
	dto, err := h.file.CreateWithMind(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, dto)
}

func (h *Handler) DeleteFile(c *gin.Context) {
	input := new(file.DeleteMind)
	if err := input.MindId.UnhashStr(c.Param("mind_id")); err != nil {
		FailErr(c, err)
		return
	}
	if err := input.FileId.UnhashStr(c.Param("file_id")); err != nil {
		FailErr(c, err)
		return
	}

	if input.FileId == 0 || input.MindId == 0 {
		Fail(c, "file_id and mind_id is required")
		return
	}

	userId := h.authMW.MustGetUserId(c)
	input.UserId = *userId
	input.DeletedBy = *userId
	err := h.file.Delete(c, input)
	if err != nil {
		FailErr(c, err)
		return
	}

	Success(c, "DELETED")
}
