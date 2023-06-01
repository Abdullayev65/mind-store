package mind

import (
	"mindstore/internal/object/dto/file"
	"mindstore/pkg/hash-types"
)

type Create struct {
	Topic     *string
	Caption   *string
	ParentId  *hash.Int
	Access    int
	HashedId  *hash.Int
	CreatedBy *hash.Int `json:"-"`
}

type Update struct {
	Id        hash.Int `json:"-"`
	Topic     *string
	Caption   *string
	ParentId  *hash.Int
	Access    *int
	HashedId  *hash.Int
	CreatedBy *hash.Int `json:"-"`
}

type List struct {
	Id       hash.Int
	Topic    *string
	Caption  *string
	ParentId *hash.Int
	Access   int
	HashedId *hash.Int
	Files    []file.List
	SubMinds []List
}

type Delete struct {
	Id        hash.Int
	DeleteBy  hash.Int
	CreatedBy hash.Int
}

type ChildrenFilter struct {
	MindId hash.Int
	UserId *hash.Int
}
