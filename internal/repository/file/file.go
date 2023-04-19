package file

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"mindstore/internal/object/dto/file"
	"mindstore/internal/object/dto/mind"
	"mindstore/internal/object/model"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
	"mindstore/pkg/repoutill"
	"mindstore/pkg/stream"
	"strconv"
	"strings"
)

type Repo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db}
}

func (r *Repo) CreateWithMind(c ctx.Ctx, files []*model.FileData, mindId hash.Int) (err error) {
	query, args, err := r.DB.BindNamed(`INSERT INTO file(path, name, hashed_id, access, size,
created_by) VALUES (:path, :name, :hashed_id, :access, :size, :created_by) RETURNING id`, &files)
	if err != nil {
		return errors.Join(errors.New("500: "), err)
	}

	tx, err := r.DB.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	err = tx.SelectContext(c, files, query, args)
	if err != nil {
		return err
	}

	mindFiles := stream.Mapper(files, func(f *model.FileData) *model.MindFile {
		return &model.MindFile{MindId: mindId, FileId: f.Id}
	})

	_, err = r.DB.NamedExecContext(c, `INSERT INTO mind_file(mind_id, file_id) 
VALUES (:mind_id, :file_id)`, &mindFiles)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) Update(c ctx.Ctx, input *mind.Update) error {
	argNum, args, setValues := 1, []any{}, []string{}
	repoutill.UpdateSetColumn(input.Topic, "topic", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.Caption, "caption", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.ParentId, "parent_id", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.Access, "access", &setValues, &args, &argNum)
	repoutill.UpdateSetColumn(input.HashedId, "hashed_id", &setValues, &args, &argNum)

	if argNum == 1 {
		return errors.New("field not found for updating")
	}
	setStr := strings.Join(setValues, " ,")
	query := fmt.Sprintf(`UPDATE mind SET %s WHERE id= %d AND deleted_at IS NULL AND created_by=%d`,
		setStr, input.Id, *input.CreatedBy)
	_, err := r.DB.ExecContext(c, query, args...)
	return err
}

func (r *Repo) GetByMindIds(c ctx.Ctx, mindIds []hash.Int) ([]file.List, error) {
	if len(mindIds) == 0 {
		return []file.List{}, nil
	}

	mindIdsStr := stream.Mapper(mindIds, func(i hash.Int) string {
		return strconv.Itoa(int(i))
	})
	mindIdsIn := strings.Join(mindIdsStr, ",")

	list := make([]file.List, 0, len(mindIds))

	err := r.DB.SelectContext(c, &list,
		fmt.Sprintf(`SELECT mf.mind_id, f.id,f.name,f.hashed_id,f.access,f.size FROM file f
INNER JOIN mind_file mf on f.id = mf.file_id
WHERE mf.deleted_at IS NULL AND mf.mind_id IN (%s) ORDER BY mind_id, f.id DESC`, mindIdsIn))

	if err != nil {
		return nil, err
	}

	return list, nil
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
