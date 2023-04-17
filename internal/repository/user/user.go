package user

import (
	"github.com/jmoiron/sqlx"
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
)

type Repo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) GetAll(c ctx.Ctx, f *user.Filter) ([]model.User, int, error) {
	//ms, query := r.filter(f)
	//
	//count, err := query.ScanAndCount(c)
	//
	//return *ms, count, err
	return nil, 0, nil
}

func (r *Repo) GetById(c ctx.Ctx, id hash.Int) (*model.User, error) {
	m := new(model.User)

	err := r.DB.GetContext(c, m, "SELECT * FROM users WHERE id= $1", id)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *Repo) Create(c ctx.Ctx, input *user.UserCreate) (*model.User, error) {
	res, err := r.DB.NamedExec(`INSERT INTO users(username, email, password, 
first_name, middle_name, last_name, birth_date) VALUES (:username, :email, :password,
:first_name, :middle_name, :last_name, :birth_date) RETURNING ID`, input)
	if err != nil {
		return nil, err
	}

	id, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	println(id)
	return nil, nil
}

func (r *Repo) Update(c ctx.Ctx, data *user.UserUpdate) (*model.User, error) {
	//m := new(model.User)
	//
	//m.Id = *data.Id
	//SetIfNotNil(m.Username, data.Username)
	//SetIfNotNil(m.Email, data.Email)
	//SetIfNotNil(m.HashPassword, data.Password)
	//SetIfNotNil(m.FirstName, data.FirstName)
	//SetIfNotNil(m.MiddleName, data.MiddleName)
	//SetIfNotNil(m.LastName, data.LastName)
	//SetIfNotNil(m.BirthDate, data.BirthDate)
	//
	//err := r.UpdateModel(c, m)
	//
	//return m, err
	return nil, nil
}

func (r *Repo) GetByUsername(c ctx.Ctx, username string) (*model.User, error) {
	return r.getBy(c, "username", username)
}

func (r *Repo) GetByEmail(c ctx.Ctx, email string) (*model.User, error) {
	return r.getBy(c, "email", email)
}

//func (r *Repo) filter(f *user.Filter) (*[]model.User, *bun.SelectQuery) {
//	ms, q := r.Filter(f.Filter)
//
//	if f.Username != nil {
//		WhereGroupAnd(q, "username = ?", *f.Username)
//	}
//	if f.Email != nil {
//		WhereGroupAnd(q, "email = ?", *f.Email)
//	}
//
//	return ms, q
//}

// specific functions

func (r *Repo) getBy(c ctx.Ctx, column string, arg any) (*model.User, error) {
	m := new(model.User)

	err := r.DB.GetContext(c, m, "SELECT * FROM users WHERE "+column+"= $1", arg)
	if err != nil {
		return nil, err
	}

	return m, nil
}
