package services

import (
	"context"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/fmiskovic/new-amz/internal/core/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestGetItemById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	t.Run("get item by valid id should return item dto", func(t *testing.T) {
		repoMock := repositories.NewItemRepositoryMock[uuid.UUID](t)
		svc := NewItemService(repoMock)

		item := entities.NewItemBuilder().
			Title("A cool book").
			Description("A cool book written by a cool author").
			Price(100).
			Build()

		repoMock.On("GetById", mock.Anything, mock.Anything).Return(*item, nil).Once()

		got, err := svc.GetById(ctx, item.ID)
		assert.Nil(t, err)
		assert.Equal(t, item.Title, got.Title)
		assert.Equal(t, item.Price, got.Price)
		repoMock.AssertCalled(t, "GetById", mock.Anything, mock.Anything)
	})
}

func TestGetItemsPage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	t.Run("get page of items should return page of item dtos", func(t *testing.T) {
		repoMock := repositories.NewItemRepositoryMock[uuid.UUID](t)
		svc := NewItemService(repoMock)

		item := entities.NewItemBuilder().
			Title("A cool book").
			Description("A cool book written by a cool author").
			Price(100).
			Build()

		page := entities.Page[entities.Item]{
			TotalPages:    1,
			TotalElements: 1,
			Elements:      []entities.Item{*item},
		}

		repoMock.On("GetPage", mock.Anything, mock.Anything).Return(page, nil).Once()

		pagable := entities.Pageable{
			Size:   10,
			Offset: 0,
		}

		got, err := svc.GetPage(ctx, pagable)
		assert.Nil(t, err)
		assert.Equal(t, 1, got.TotalElements)
		assert.Equal(t, 1, len(got.Elements))
		assert.Equal(t, item.Title, got.Elements[0].Title)
		repoMock.AssertCalled(t, "GetPage", mock.Anything, mock.Anything)
	})
}
