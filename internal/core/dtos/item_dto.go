package dtos

import (
	"github.com/fmiskovic/new-amz/internal/core/entities"
	"time"
)

type ItemDto struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"name"`
	Description string    `json:"Description"`
	Price       float32   `json:"Price"`
}

func ToItemDto(item entities.Item) ItemDto {
	return ItemDto{
		ID:          item.ID.String(),
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		Title:       item.Title,
		Description: item.Description,
		Price:       item.Price,
	}
}

// ToPageItemDto converts Item entities Page into a Item DTO Page.
func ToPageItemDto(page entities.Page[entities.Item]) entities.Page[ItemDto] {
	dtos := make([]ItemDto, len(page.Elements))
	for i, item := range page.Elements {
		dtos[i] = ToItemDto(item)
	}
	return entities.Page[ItemDto]{
		TotalPages:    page.TotalPages,
		TotalElements: page.TotalElements,
		Elements:      dtos,
	}
}
