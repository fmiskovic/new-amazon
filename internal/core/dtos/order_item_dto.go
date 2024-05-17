package dtos

import (
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"github.com/google/uuid"
	"time"
)

type OrderItemDto struct {
	OrderID   string    `json:"order_id"`
	ItemID    string    `json:"item_id"`
	Quantity  int       `json:"quantity" validate:"required,gte=1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToOrderItemDto(orderItem entities.OrderItem) OrderItemDto {
	return OrderItemDto{
		OrderID:   orderItem.OrderID.String(),
		ItemID:    orderItem.ItemID.String(),
		Quantity:  orderItem.Quantity,
		CreatedAt: orderItem.CreatedAt,
		UpdatedAt: orderItem.UpdatedAt,
	}
}

func ToOrderItemEntity(dto OrderItemDto) (*entities.OrderItem, error) {
	itemId, err := uuid.Parse(dto.ItemID)
	if err != nil {
		return nil, err
	}

	return &entities.OrderItem{
		ItemID:   itemId,
		Quantity: dto.Quantity,
	}, nil
}

func ToOrderItemEntities(dtos []OrderItemDto) ([]*entities.OrderItem, error) {
	items := make([]*entities.OrderItem, len(dtos))
	for i, dto := range dtos {
		entity, err := ToOrderItemEntity(dto)
		if err != nil {
			return nil, err
		}
		items[i] = entity
	}
	return items, nil
}
