package file

import (
	"github.com/gin-gonic/gin"
	file_srvc "mindstore/internal/service/file"
	"mindstore/pkg/hash-types"
	. "mindstore/pkg/response"
)

type Handler struct {
	authMW AuthMW
	file   *file_srvc.Service
}

func New(authMW AuthMW, file *file_srvc.Service) *Handler {
	return &Handler{authMW, file}
}

func (h *Handler) GetFile(c *gin.Context) {
	userId := *h.authMW.MustGetUserId(c)

	var fileId hash.Int
	err := fileId.UnhashStr(c.Param("id"))
	if err != nil {
		FailErr(c, err)
		return
	}

	path, err := h.file.GetPathById(c, fileId, userId)
	if err != nil {
		FailErr(c, err)
		return
	}

	c.File(path)
}
