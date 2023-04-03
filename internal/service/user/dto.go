package user

import (
	"mindstore/internal/basic"
	"time"
)

type UserCreate struct {
	Username   *string    `json:"username"`
	Email      *string    `json:"email"`
	Password   *string    `json:"password"`
	FirstName  *string    `json:"first_name"`
	MiddleName *string    `json:"middle_name"`
	LastName   *string    `json:"last_name"`
	BirthDate  *time.Time `json:"birth_date"`
}

type UserUpdate struct {
	Id         *int       `json:"-"`
	Username   *string    `json:"username"`
	Email      *string    `json:"email"`
	Password   *string    `json:"password"`
	FirstName  *string    `json:"first_name"`
	MiddleName *string    `json:"middle_name"`
	LastName   *string    `json:"last_name"`
	BirthDate  *time.Time `json:"birth_date"`
}

type UserDetail struct {
	Id         *int       `json:"id"`
	Username   *string    `json:"username"`
	Email      *string    `json:"email"`
	MindId     *int       `json:"mind_id"`
	FirstName  *string    `json:"first_name"`
	MiddleName *string    `json:"middle_name"`
	LastName   *string    `json:"last_name"`
	BirthDate  *time.Time `json:"birth_date"`
}

type Filter struct {
	basic.Filter

	Username *string
	Email    *string
}
