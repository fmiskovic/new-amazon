package repositories

import (
	"context"
	"database/sql"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"sync"
)

type OrderRepository struct {
	bunDb *bun.DB
	mutex sync.RWMutex
}

func NewOrderRepository(db *bun.DB) *OrderRepository {
	return &OrderRepository{db, sync.RWMutex{}}
}

func (repo *OrderRepository) GetById(ctx context.Context, id uuid.UUID) (entities.Order, error) {
	var order = new(entities.Order)

	err := repo.bunDb.NewSelect().
		Model(order).
		Relation("OrderItems").
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return entities.Order{}, err
	}

	return *order, nil
}

func (repo *OrderRepository) Search(ctx context.Context, accountId uuid.UUID, p entities.Pageable) (entities.Page[entities.Order], error) {
	var orders []entities.Order

	count, err := repo.bunDb.NewSelect().
		Model(&orders).
		Relation("OrderItems").
		Where("account_id = ?", accountId).
		Limit(p.Size).
		Offset(p.Offset).
		Order(entities.StringifyOrders(p.Sort)...).
		ScanAndCount(ctx)

	totalPages := 0
	if count != 0 && p.Size != 0 {
		totalPages = (len(orders) + p.Size - 1) / p.Size
	}

	return entities.Page[entities.Order]{
		TotalPages:    totalPages,
		TotalElements: count,
		Elements:      orders,
	}, err
}

func (repo *OrderRepository) Create(ctx context.Context, order *entities.Order) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	if order == nil {
		return ErrNilEntity
	}

	return repo.bunDb.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		ok, err := tx.NewSelect().Model(&entities.Account{}).Where("id = ?", order.AccountID).Exists(ctx)
		if err != nil {
			return err
		}
		if !ok {
			return ErrNotFound
		}

		_, err = tx.NewInsert().Model(order).Exec(ctx)
		if err != nil {
			return err
		}

		if order.OrderItems != nil {
			for _, item := range order.OrderItems {
				if item == nil {
					continue
				}
				item.OrderID = order.ID
				_, err = tx.NewInsert().Model(item).Exec(ctx)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}
