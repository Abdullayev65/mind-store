package mind

import (
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
	Id       hash.Int `json:"-"`
	Topic    *string
	Caption  *string
	ParentId *hash.Int
	Access   int
	HashedId *hash.Int
	UserId   *hash.Int `json:"-"`
}

type Detail struct {
	Id       hash.Int
	Topic    *string
	Caption  *string
	ParentId *int
	Access   int
	HashedId *int
	Files    []string
}
