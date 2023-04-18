package user

import (
	"mindstore/internal/object/dto"
	"mindstore/pkg/hash-types"
	"time"
)

type UserCreate struct {
	Username     *string
	Email        *string
	Password     *string
	FirstName    *string
	MiddleName   *string
	LastName     *string
	BirthDate    *time.Time `json:"-"`
	BirthDateStr *string    `json:"birth_date"`
}

type UserUpdate struct {
	Id           hash.Int `json:"-" form:"-"`
	Username     *string
	Email        *string
	Password     *string
	FirstName    *string
	MiddleName   *string
	LastName     *string
	BirthDate    *time.Time `json:"-"`
	BirthDateStr *string    `json:"birth_date"`
}

type UserDetail struct {
	Id           hash.Int
	Username     *string
	Email        *string
	MindId       *hash.Int
	FirstName    *string
	MiddleName   *string
	LastName     *string
	BirthDate    *time.Time `json:"-"`
	BirthDateStr *string    `json:"birth_date"`
}

type Filter struct {
	dto.Filter

	Username *string
	Email    *string
}
