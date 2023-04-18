package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"mindstore/internal/object/dto/user"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
	"mindstore/pkg/repoutill"
	"strings"
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
	res, err := r.DB.NamedExecContext(c, `INSERT INTO users(username, email, password, 
first_name, middle_name, last_name, birth_date) VALUES (:username, :email, :password,
:first_name, :middle_name, :last_name, :birth_date) RETURNING id`, input)
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
func (r *Repo) CreateWithMind(c ctx.Ctx, input *user.UserCreate) (hash.Int, error) {
	txx, err := r.DB.BeginTxx(c, &sql.TxOptions{})
	if err != nil {
		return 0, err
	}
	defer txx.Commit()

	res, err := txx.NamedExecContext(c, `INSERT INTO users(username, email, password, 
first_name, middle_name, last_name, birth_date) VALUES (:username, :email, :password,
:first_name, :middle_name, :last_name, :birth_date) RETURNING id`, input)
	if err != nil {
		txx.Rollback()
		return 0, err
	}

	id, err := res.RowsAffected()
	if err != nil {
		txx.Rollback()
		return 0, err
	}

	res, err = txx.ExecContext(c, `INSERT INTO mind(topic, created_by) 
VALUES ($1, $2)`, input.Username, id)
	if err != nil {
		txx.Rollback()
		return 0, err
	}

	return hash.Int(id), nil
}

func (r *Repo) Update(c ctx.Ctx, input *user.UserUpdate) error {
	argNum, args, setValues := 1, []any{}, []string{}
	repoutill.UpdateSetColumn(input.Username, "username", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.Email, "email", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.Password, "password", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.FirstName, "first_name", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.MiddleName, "middle_name", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.LastName, "last_name", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.BirthDate, "birth_date", &setValues, &args, &argNum)

	if argNum == 1 {
		return errors.New("field not found for updating")
	}
	setStr := strings.Join(setValues, " ,")
	query := fmt.Sprintf(`UPDATE users u SET %s WHERE id= %d AND deleted_at IS NULL`,
		setStr, input.Id)
	_, err := r.DB.ExecContext(c, query, args...)
	return err
}

func (r *Repo) GetByUsername(c ctx.Ctx, username string) (*model.User, error) {
	return r.getBy(c, "username", username)
}

func (r *Repo) GetByEmail(c ctx.Ctx, email string) (*model.User, error) {
	return r.getBy(c, "email", email)
}

func (r *Repo) DetailById(c ctx.Ctx, id *hash.Int) (*user.UserDetail, error) {
	o := new(user.UserDetail)

	err := r.DB.GetContext(c, o,
		`SELECT id, username, email, mind_id, first_name, 
middle_name, last_name, birth_date FROM users WHERE id=$1`, id)

	if err != nil {
		return nil, err
	}

	return o, nil
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

	err := r.DB.GetContext(c, m, "SELECT * FROM users WHERE deleted_at IS NULL AND "+
		column+"= $1, ", arg)
	if err != nil {
		return nil, err
	}

	return m, nil
}
