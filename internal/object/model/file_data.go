package model

import (
	"mindstore/internal/object/model/submodel"
)

type FileData struct {
	submodel.BasicModel

	Path     string
	Name     string
	HashedId *int
	Access   int
	Size     int64
}
