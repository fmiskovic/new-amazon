package services

import (
	"context"
	"errors"
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/fmiskovic/new-amz/internal/core/repositories"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateOrder(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	t.Run("create order with valid data should return no error", func(t *testing.T) {
		repoMock := repositories.NewOrderRepositoryMock[uuid.UUID](t)
		svc := NewOrderService(repoMock)

		orderIdMock := uuid.NewString()

		items := make([]dtos.OrderItemDto, 2)
		items[0] = dtos.OrderItemDto{
			ItemID:   uuid.NewString(),
			OrderID:  orderIdMock,
			Quantity: 2,
		}
		items[1] = dtos.OrderItemDto{
			ItemID:   uuid.NewString(),
			OrderID:  orderIdMock,
			Quantity: 1,
		}

		cmd := dtos.CreateOrderCommand{
			AccountID: uuid.NewString(),
			Items:     items,
		}

		repoMock.On("Create", mock.Anything, mock.Anything).Return(nil).Once()

		got, err := svc.Create(ctx, cmd)
		assert.Nil(t, err)
		assert.Equal(t, cmd.AccountID, got.AccountID)
		repoMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
	})

	t.Run("create order with invalid account id should return error", func(t *testing.T) {
		repoMock := repositories.NewOrderRepositoryMock[uuid.UUID](t)
		svc := NewOrderService(repoMock)

		cmd := dtos.CreateOrderCommand{
			AccountID: "invalid-uuid",
		}

		_, err := svc.Create(ctx, cmd)
		assert.NotNil(t, err)
		repoMock.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	})

	t.Run("create order with invalid item id should return error", func(t *testing.T) {
		repoMock := repositories.NewOrderRepositoryMock[uuid.UUID](t)
		svc := NewOrderService(repoMock)

		orderIdMock := uuid.NewString()

		items := make([]dtos.OrderItemDto, 2)
		items[0] = dtos.OrderItemDto{
			ItemID:   "invalid-uuid",
			OrderID:  orderIdMock,
			Quantity: 2,
		}

		cmd := dtos.CreateOrderCommand{
			AccountID: uuid.NewString(),
			Items:     items,
		}

		_, err := svc.Create(ctx, cmd)
		assert.NotNil(t, err)
		repoMock.AssertNotCalled(t, "Create", mock.Anything, mock.Anything)
	})

	t.Run("create order fails should return error", func(t *testing.T) {
		repoMock := repositories.NewOrderRepositoryMock[uuid.UUID](t)
		svc := NewOrderService(repoMock)

		orderIdMock := uuid.NewString()

		items := make([]dtos.OrderItemDto, 2)
		items[0] = dtos.OrderItemDto{
			ItemID:   uuid.NewString(),
			OrderID:  orderIdMock,
			Quantity: 2,
		}
		items[1] = dtos.OrderItemDto{
			ItemID:   uuid.NewString(),
			OrderID:  orderIdMock,
			Quantity: 1,
		}

		cmd := dtos.CreateOrderCommand{
			AccountID: uuid.NewString(),
			Items:     items,
		}

		repoMock.On("Create", mock.Anything, mock.Anything).Return(errors.New("unexpected error")).Once()

		_, err := svc.Create(ctx, cmd)
		assert.NotNil(t, err)
		repoMock.AssertCalled(t, "Create", mock.Anything, mock.Anything)
	})
}

func TestSearchAccountOrders(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	t.Run("search orders with valid account id should return page of orders", func(t *testing.T) {
		repoMock := repositories.NewOrderRepositoryMock[uuid.UUID](t)
		svc := NewOrderService(repoMock)

		mockItems := make([]*entities.OrderItem, 1)
		mockItems[0] = &entities.OrderItem{
			ItemID:   uuid.New(),
			OrderID:  uuid.New(),
			Quantity: 2,
		}

		mockOrder := entities.NewOrderBuilder().
			AccountID(uuid.New()).
			OrderItems(mockItems).
			Build()

		mockPage := entities.Page[entities.Order]{
			TotalPages:    1,
			TotalElements: 1,
			Elements:      []entities.Order{*mockOrder},
		}
		repoMock.On("Search", mock.Anything, mock.Anything, mock.Anything).Return(mockPage, nil).Once()

		accountId := uuid.New()
		pageRequest := entities.Pageable{Size: 10, Offset: 0}

		filter := dtos.OrderFilter{
			AccountID:   accountId,
			PageRequest: pageRequest,
		}

		page, err := svc.Search(ctx, filter)
		assert.Nil(t, err)
		assert.Equal(t, 1, page.TotalElements)
		assert.Equal(t, mockOrder.AccountID.String(), page.Elements[0].AccountID)
		repoMock.AssertCalled(t, "Search", mock.Anything, accountId, pageRequest)
	})

	t.Run("search orders returns unexpected error", func(t *testing.T) {
		repoMock := repositories.NewOrderRepositoryMock[uuid.UUID](t)
		svc := NewOrderService(repoMock)

		mockPage := entities.Page[entities.Order]{
			TotalPages:    1,
			TotalElements: 1,
			Elements:      []entities.Order{},
		}

		repoMock.On("Search", mock.Anything, mock.Anything, mock.Anything).
			Return(mockPage, errors.New("unexpected error")).Once()

		accountId := uuid.New()
		pageRequest := entities.Pageable{Size: 10, Offset: 0}

		filter := dtos.OrderFilter{
			AccountID:   accountId,
			PageRequest: pageRequest,
		}

		_, err := svc.Search(ctx, filter)
		assert.NotNil(t, err)
		repoMock.AssertCalled(t, "Search", mock.Anything, accountId, pageRequest)
	})
}

func TestGetOrderById(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	t.Run("get order by id should return order", func(t *testing.T) {
		repoMock := repositories.NewOrderRepositoryMock[uuid.UUID](t)
		svc := NewOrderService(repoMock)

		mockItems := make([]*entities.OrderItem, 1)
		mockItems[0] = &entities.OrderItem{
			ItemID:   uuid.New(),
			OrderID:  uuid.New(),
			Quantity: 2,
		}

		mockOrder := entities.NewOrderBuilder().
			AccountID(uuid.New()).
			OrderItems(mockItems).
			Build()

		repoMock.On("GetById", mock.Anything, mock.Anything).Return(*mockOrder, nil).Once()

		got, err := svc.GetById(ctx, mockOrder.ID)
		assert.Nil(t, err)
		assert.Equal(t, mockOrder.AccountID.String(), got.AccountID)
		repoMock.AssertCalled(t, "GetById", mock.Anything, mock.Anything)
	})

	t.Run("get order by id returns unexpected error", func(t *testing.T) {
		repoMock := repositories.NewOrderRepositoryMock[uuid.UUID](t)
		svc := NewOrderService(repoMock)

		repoMock.On("GetById", mock.Anything, mock.Anything).
			Return(entities.Order{}, errors.New("unexpected error")).Once()

		_, err := svc.GetById(ctx, uuid.New())
		assert.NotNil(t, err)
		repoMock.AssertCalled(t, "GetById", mock.Anything, mock.Anything)
	})
}
