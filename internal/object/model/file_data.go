package model

import (
	"github.com/uptrace/bun"
	"mindstore/internal/object/model/submodel"
)

type FileData struct {
	bun.BaseModel
	submodel.BasicModel

	Path     string
	HashedId *int
	Access   int
	Size     int
}
