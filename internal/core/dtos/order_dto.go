package dtos

import (
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/google/uuid"
	"time"
)

type OrderDto struct {
	ID           string         `json:"id"`
	AccountID    string         `json:"account_id"`
	AccountEmail string         `json:"account_email"`
	Items        []OrderItemDto `json:"items"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}

func ToOrderDto(order entities.Order) OrderDto {
	items := make([]OrderItemDto, len(order.OrderItems))
	for i, item := range order.OrderItems {
		if item == nil {
			continue
		}
		items[i] = ToOrderItemDto(*item)
	}
	return OrderDto{
		ID:           order.ID.String(),
		CreatedAt:    order.CreatedAt,
		AccountID:    order.AccountID.String(),
		AccountEmail: order.Account.Email,
		Items:        items,
	}
}

type CreateOrderCommand struct {
	AccountID string         `json:"account_id"`
	Items     []OrderItemDto `json:"items"`
}

type CreateOrderAnswer struct {
	OrderDto
}

type OrderFilter struct {
	AccountID   uuid.UUID
	PageRequest entities.Pageable
}

// ToPageOrderDto converts Order entities Page into a Order DTO Page.
func ToPageOrderDto(page entities.Page[entities.Order]) entities.Page[OrderDto] {
	dtos := make([]OrderDto, len(page.Elements))
	for i, order := range page.Elements {
		dtos[i] = ToOrderDto(order)
	}
	return entities.Page[OrderDto]{
		TotalPages:    page.TotalPages,
		TotalElements: page.TotalElements,
		Elements:      dtos,
	}
}
