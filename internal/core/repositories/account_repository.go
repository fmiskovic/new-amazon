package repositories

import (
	"github.com/fmiskovic/new-amz/internal/core/entities"

	"context"
)

// AccountRepository is a secondary port for account operations.
type AccountRepository[ID any] interface {
	GetById(ctx context.Context, id ID) (entities.Account, error)
	GetByEmail(ctx context.Context, email string) (entities.Account, error)
	Create(ctx context.Context, account *entities.Account) error
}
