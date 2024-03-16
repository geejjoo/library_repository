package dao

import (
	"context"
	"github.com/geejjoo/library_repository/internal/db/tabler"
)

type Adapter interface {
	Create(ctx context.Context, entity tabler.Tabler) error
	Update(ctx context.Context, entity tabler.Tabler, condition Condition) error
	List(ctx context.Context, dest interface{}, entity tabler.Tabler, condition Condition) error
	Delete(ctx context.Context, entity tabler.Tabler, condition Condition) error
}
