package repositories

import (
	"context"
	"github.com/fmiskovic/new-amz/internal/core/entities"
)

// OrderRepository is an interface for interacting with the order repository.
type OrderRepository[ID any] interface {
	GetById(ctx context.Context, id ID) (entities.Order, error)
	Search(ctx context.Context, accountId ID, pageRequest entities.Pageable) (entities.Page[entities.Order], error)
	Create(ctx context.Context, order *entities.Order) error
}
