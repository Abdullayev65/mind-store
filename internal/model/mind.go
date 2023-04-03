package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Mind struct {
	bun.BaseModel
	BasicModel

	Name        *string
	Description *string
	ParentId    *int
	Access      *int
	hashed      bool
	UpdatedAt   time.Time `bun:",default:now(),notnull"`
	UpdatedBy   int
}

// access
// 33 - private
// 66 - friends - beta
// 99 - public
