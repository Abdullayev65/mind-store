package repository

import (
	"mindstore/internal/db"
	"mindstore/internal/repository/user"
)

var User = new(user.Repo)

func init() {
	*User = *user.New(db.Sqlx)
}
