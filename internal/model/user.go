package model

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel
	BasicModel

	Username     *string    `bun:",uniq"`
	Email        *string    `bun:",uniq"`
	MindId       *int       `bun:"mind_id"`
	HashPassword *string    `bun:"hash_password"`
	FirstName    *string    `bun:"first_name"`
	MiddleName   *string    `bun:"middle_name"`
	LastName     *string    `bun:"last_name"`
	BirthDate    *time.Time `bun:"birth_date"`
}
