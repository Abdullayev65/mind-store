package mind

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"mindstore/internal/object/dto/mind"
	"mindstore/internal/object/dto/user"
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

func (r *Repo) Create(c ctx.Ctx, input *mind.Create) (*hash.Int, error) {
	res, err := r.DB.NamedQueryContext(c, `INSERT INTO mind(topic, caption, parent_id, access, 
hashed_id, created_by) VALUES (:topic, :caption, :parent_id, :access, :hashed_id, :created_by) 
RETURNING id`, input)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	if !res.Next() {
		return nil, errors.New("500")
	}
	id := new(hash.Int)
	err = res.Scan(id)
	if err != nil {
		return nil, err
	}

	return id, nil
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
	query := fmt.Sprintf(`UPDATE users SET %s WHERE id= %d AND deleted_at IS NULL`,
		setStr, input.Id)
	_, err := r.DB.ExecContext(c, query, args...)
	return err
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

func (r *Repo) Delete(c ctx.Ctx, userId hash.Int, deletedBy hash.Int) error {
	_, err := r.DB.ExecContext(c, `UPDATE users SET deleted_at = now(), deleted_by = $1
	 WHERE id = $2`, deletedBy, userId)
	if err != nil {
		return err
	}

	return err
}

//func (r *Repo) filter(f *user.Filter) (*[]model.user, *bun.SelectQuery) {
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

//func (r *Repo) GetAll(c ctx.Ctx, f *user.Filter) ([]model.User, int, error) {
//	//ms, query := r.filter(f)
//	//
//	//count, err := query.ScanAndCount(c)
//	//
//	//return *ms, count, err
//	return nil, 0, nil
//}
//
//func (r *Repo) GetById(c ctx.Ctx, id hash.Int) (*model.User, error) {
//	m := new(model.Mind)
//
//	err := r.DB.GetContext(c, m, "SELECT * FROM users WHERE id= $1", id)
//	if err != nil {
//		return nil, err
//	}
//
//	return m, nil
//}
