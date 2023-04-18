package mind

import (
	"mindstore/internal/object/dto/mind"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
)

type Mind interface {
	Create(c ctx.Ctx, input *mind.Create) (*hash.Int, error)
}
