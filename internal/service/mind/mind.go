package mind

import (
	"errors"
	"mindstore/internal/object/dto/mind"
	"mindstore/pkg/ctx"
	"mindstore/pkg/hash-types"
)

type Service struct {
	mind Mind
}

func New(mind Mind) *Service {
	return &Service{mind}
}

func (s *Service) CreateMind(c ctx.Ctx, input *mind.Create) (*hash.Int, error) {
	var errStr string
	switch {
	case input.CreatedBy == nil:
		errStr = "owner not found"
	case input.Topic == nil:
		errStr = "topic can't be null"
	case input.ParentId == nil:
		errStr = "parent of mind not given"
	}
	if errStr != "" {
		return nil, errors.New(errStr)
	}

	if input.Access != 33 && input.Access != 99 {
		input.Access = 33
	}

	return s.mind.Create(c, input)
}
