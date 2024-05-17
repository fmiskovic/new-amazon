package repositories

import (
	"context"
	"github.com/fmiskovic/new-amz/internal/core/entities"
)

// ItemRepository is a secondary port for item operations.
type ItemRepository[ID any] interface {
	GetById(ctx context.Context, id ID) (entities.Item, error)
	GetPage(ctx context.Context, p entities.Pageable) (entities.Page[entities.Item], error)
}
