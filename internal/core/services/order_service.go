package services

import (
	"context"
	"github.com/fmiskovic/new-amz/internal/core/dtos"
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/fmiskovic/new-amz/internal/core/repositories"
	"github.com/google/uuid"
)

// OrderService represents business logic related to entities.Order.
type OrderService struct {
	repo repositories.OrderRepository[uuid.UUID]
}

// NewOrderService instantiates new OrderService.
func NewOrderService(repo repositories.OrderRepository[uuid.UUID]) OrderService {
	return OrderService{repo: repo}
}

// GetById returns existing order.
func (s OrderService) GetById(ctx context.Context, id uuid.UUID) (dtos.OrderDto, error) {
	order, err := s.repo.GetById(ctx, id)
	if err != nil {
		return dtos.OrderDto{}, err
	}

	return dtos.ToOrderDto(order), nil
}

// Search returns page of orders created by specified account.
func (s OrderService) Search(ctx context.Context, filter dtos.OrderFilter) (entities.Page[dtos.OrderDto], error) {
	orders, err := s.repo.Search(ctx, filter.AccountID, filter.PageRequest)
	if err != nil {
		return entities.Page[dtos.OrderDto]{}, err
	}

	return dtos.ToPageOrderDto(orders), nil
}

// Create creates new order.
func (s OrderService) Create(ctx context.Context, cmd dtos.CreateOrderCommand) (dtos.CreateOrderAnswer, error) {
	accountId, err := uuid.Parse(cmd.AccountID)
	if err != nil {
		return dtos.CreateOrderAnswer{}, newError("invalid account id", err)
	}

	orderItems, err := dtos.ToOrderItemEntities(cmd.Items)
	if err != nil {
		return dtos.CreateOrderAnswer{}, newError("invalid order items", err)
	}

	order := entities.NewOrderBuilder().
		AccountID(accountId).
		OrderItems(orderItems).
		Build()

	if err := s.repo.Create(ctx, order); err != nil {
		return dtos.CreateOrderAnswer{}, err
	}

	return dtos.CreateOrderAnswer{OrderDto: dtos.ToOrderDto(*order)}, nil
}
