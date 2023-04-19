package file

import (
	"mime/multipart"
	"mindstore/internal/object/dto/file"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
)

type File interface {
	CreateWithMind(c ctx.Ctx, files []*model.FileData, mindId hash.Int) error
	GetByMindIds(c ctx.Ctx, mindIds []hash.Int) ([]file.List, error)
}

type SysFile interface {
	MultipleUploadFile(files []*multipart.FileHeader, folder string) (fds []*model.FileData, err error)
}
