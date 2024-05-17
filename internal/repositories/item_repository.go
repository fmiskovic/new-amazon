package repositories

import (
	"context"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ItemRepository is the implementation of core repositories.ItemRepository interface.
type ItemRepository struct {
	bunDb *bun.DB
}

// NewItemRepository instantiates new ItemRepository.
func NewItemRepository(db *bun.DB) ItemRepository {
	return ItemRepository{db}
}

// GetById returns item by specified id.
func (repo ItemRepository) GetById(ctx context.Context, id uuid.UUID) (entities.Item, error) {
	item := new(entities.Item)

	err := repo.bunDb.NewSelect().Model(item).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return *item, err
	}

	return *item, nil
}

// GetPage respond with a page of items.
func (repo ItemRepository) GetPage(ctx context.Context, p entities.Pageable) (entities.Page[entities.Item], error) {
	var items []entities.Item
	count, err := repo.bunDb.NewSelect().
		Model(&items).
		Limit(p.Size).
		Offset(p.Offset).
		Order(entities.StringifyOrders(p.Sort)...).
		ScanAndCount(ctx)

	totalPages := 0
	if count != 0 && p.Size != 0 {
		totalPages = (len(items) + p.Size - 1) / p.Size
	}

	return entities.Page[entities.Item]{
		TotalPages:    totalPages,
		TotalElements: count,
		Elements:      items,
	}, err
}
