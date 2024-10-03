package storage

import (
	"context"

	"gitlab.com/golight/orm/utils"
)

type SQLAdapter interface {
	Create(ctx context.Context, entity utils.Tabler, opts ...interface{}) error
	Upsert(ctx context.Context, entities []utils.Tabler, opts ...interface{}) error
	GetCount(ctx context.Context, entity utils.Tabler, condition utils.Condition, opts ...interface{}) (uint64, error)
	List(ctx context.Context, dest interface{}, tableName string, condition utils.Condition, opts ...interface{}) error
	Update(ctx context.Context, entity utils.Tabler, condition utils.Condition, operation string, opts ...interface{}) error
}
