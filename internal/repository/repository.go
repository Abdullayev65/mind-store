package repository

import (
	"mindstore/internal/db"
	"mindstore/internal/repository/mind"
	"mindstore/internal/repository/user"
)

var User = new(user.Repo)
var Mind = new(mind.Repo)

func init() {
	*User = *user.New(db.Sqlx)
	*Mind = *mind.New(db.Sqlx)
}
