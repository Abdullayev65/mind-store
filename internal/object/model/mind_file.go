package model

import "mindstore/internal/object/model/submodel"

type MindFile struct {
	MindId int
	FileId int
	submodel.BasicTimeModel
}
