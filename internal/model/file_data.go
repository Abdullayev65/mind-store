package model

import "github.com/uptrace/bun"

type FileData struct {
	bun.BaseModel
	BasicModel

	Path   string
	Access int
}
