package services

import (
	"context"
	"fmt"
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/fmiskovic/new-amz/internal/core/repositories"
	"github.com/google/uuid"
)

// ItemService represents business logic related to entities.Item.
type ItemService struct {
	repo repositories.ItemRepository[uuid.UUID]
}

// NewItemService instantiates new ItemService.
func NewItemService(repo repositories.ItemRepository[uuid.UUID]) ItemService {
	return ItemService{repo}
}

// GetById returns existing item by id.
func (s ItemService) GetById(ctx context.Context, id uuid.UUID) (dtos.ItemDto, error) {
	item, err := s.repo.GetById(ctx, id)
	if err != nil {
		return dtos.ItemDto{}, newError(fmt.Sprintf("failed to get item by id: %s", id.String()), err)
	}
	return dtos.ToItemDto(item), nil
}

// GetPage returns page of items.
func (s ItemService) GetPage(ctx context.Context, p entities.Pageable) (entities.Page[dtos.ItemDto], error) {
	page, err := s.repo.GetPage(ctx, p)
	if err != nil {
		return entities.Page[dtos.ItemDto]{}, newError("failed to get page of items", err)
	}
	return dtos.ToPageItemDto(page), nil
}
