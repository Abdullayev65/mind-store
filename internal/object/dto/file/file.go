package file

import "mindstore/pkg/hash-types"

type List struct {
	Id       hash.Int
	MindId   hash.Int
	Url      string
	Name     string
	HashedId *int
	Access   int
	Size     int
}
