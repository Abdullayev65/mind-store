package user

import (
	"context"
	"github.com/uptrace/bun"
	"mindstore/internal/basic"
	. "mindstore/internal/basic/repo"
	"mindstore/internal/model"
	"mindstore/internal/service/user"
)

type Repo struct {
	DB *bun.DB
	*basic.Repository[model.User]
}

func (r *Repo) GetAll(c context.Context, f *user.Filter) ([]model.User, int, error) {
	ms, query := r.filter(f)

	count, err := query.ScanAndCount(c)

	return *ms, count, err
}

func (r *Repo) FindAll(c context.Context, f *user.Filter) ([]model.User, error) {
	ms, query := r.filter(f)

	err := query.Scan(c)

	return *ms, err
}

func (r *Repo) Create(c context.Context, data *user.UserCreate) (*model.User, error) {
	m := new(model.User)

	m.Username = data.Username
	m.Email = data.Email
	m.HashPassword = data.Password
	m.FirstName = data.FirstName
	m.MiddleName = data.MiddleName
	m.LastName = data.LastName
	m.BirthDate = data.BirthDate

	err := r.InsertModel(c, m)

	return m, err
}

func (r *Repo) Update(c context.Context, data *user.UserUpdate) (*model.User, error) {
	m := new(model.User)

	m.Id = *data.Id
	SetIfNotNil(m.Username, data.Username)
	SetIfNotNil(m.Email, data.Email)
	SetIfNotNil(m.HashPassword, data.Password)
	SetIfNotNil(m.FirstName, data.FirstName)
	SetIfNotNil(m.MiddleName, data.MiddleName)
	SetIfNotNil(m.LastName, data.LastName)
	SetIfNotNil(m.BirthDate, data.BirthDate)

	err := r.UpdateModel(c, m)

	return m, err
}

func (r *Repo) GetByUsername(c context.Context, username string) (*model.User, error) {
	m := new(model.User)
	err := r.DB.NewSelect().Model(m).
		Where("username = ?", username).Scan(c)

	return m, err
}

func (r *Repo) GetByEmail(c context.Context, email string) (*model.User, error) {
	m := new(model.User)
	err := r.DB.NewSelect().Model(m).
		Where("email = ?", email).Scan(c)

	return m, err
}

func (r *Repo) filter(f *user.Filter) (*[]model.User, *bun.SelectQuery) {
	ms, q := r.Filter(f.Filter)

	if f.Username != nil {
		WhereGroupAnd(q, "username = ?", *f.Username)
	}
	if f.Email != nil {
		WhereGroupAnd(q, "email = ?", *f.Email)
	}

	return ms, q
}
