package basic

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type Repository[M any] struct {
	DB *bun.DB
}

func (r *Repository[M]) InsertModel(c context.Context, m *M) error {
	_, err := r.DB.NewInsert().Model(m).Exec(c)
	return err
}

func (r *Repository[M]) UpdateModel(c context.Context, m *M) error {
	_, err := r.DB.NewUpdate().Model(m).
		WherePK().Exec(c)
	return err
}

func (r *Repository[M]) GetById(c context.Context, id int) (*M, error) {
	m := new(M)
	err := r.DB.NewSelect().Model(m).
		Where("id = ?", id).Scan(c)
	return m, err
}

func (r *Repository[M]) Delete(c context.Context, id, deletedBy int) error {
	var m *M
	_, err := r.DB.NewUpdate().
		Model(m).
		Set("deleted_at = ?", time.Now()).
		Set("deleted_by = ?", deletedBy).
		Where("id = ?", id).Exec(c)

	return err
}

func (r *Repository[M]) Filter(f Filter) (*[]M, *bun.SelectQuery) {
	ms := make([]M, 0, 3)
	q := r.DB.NewSelect().Model(&ms)

	if f.Limit != nil {
		q.Limit(*f.Limit)
	}
	if f.Offset != nil {
		q.Offset(*f.Offset)
	}
	if f.Order != nil {
		q.Order(*f.Order)
	} else {
		q.Order("id desc")
	}
	if f.Deleted {
		q.WhereDeleted()
	}
	if f.WithDeleted {
		q.WhereAllWithDeleted()
	}

	return &ms, q
}
