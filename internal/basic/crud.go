package basic

import "context"

type CRUD[M, C, U, F any] interface {
	GetAll(ctx context.Context, filter *F) ([]M, int, error)
	FindAll(ctx context.Context, filter *F) ([]M, error)
	GetById(ctx context.Context, id int) (*M, error)
	Create(ctx context.Context, create *C) (*M, error)
	Update(ctx context.Context, update *U) (*M, error)
	Delete(ctx context.Context, id int, deletedBy int) error
}
