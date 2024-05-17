package repositories

import (
	"context"
	"database/sql"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// AccountRepository is the implementation of core repositories.AccountRepository interface.
type AccountRepository struct {
	db *bun.DB
}

// NewAccountRepository instantiates new AccountRepository.
func NewAccountRepository(db *bun.DB) AccountRepository {
	return AccountRepository{db}
}

// GetById returns account by specified id.
func (repo AccountRepository) GetById(ctx context.Context, id uuid.UUID) (entities.Account, error) {
	var acc = new(entities.Account)

	err := repo.db.NewSelect().Model(acc).Where("? = ?", bun.Ident("id"), id).Scan(ctx)
	if err != nil {
		return *acc, err
	}

	return *acc, nil
}

// GetByEmail returns account by email.
func (repo AccountRepository) GetByEmail(ctx context.Context, email string) (entities.Account, error) {
	var u = new(entities.Account)

	err := repo.db.NewSelect().
		Model(u).
		Where("email = ?", email).
		Scan(ctx)

	if err != nil {
		return entities.Account{}, err
	}

	return *u, nil
}

// Create persists new account entity.
func (repo AccountRepository) Create(ctx context.Context, u *entities.Account) error {
	if u == nil {
		return ErrNilEntity
	}

	return repo.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(u).Exec(ctx)
		return err
	})
}
