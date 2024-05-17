package entities

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

// Item will store information about each item in the store.
type Item struct {
	bun.BaseModel `bun:"table:items,alias:i"`

	Entity
	Title       string  `bun:"title,nullzero"`
	Description string  `bun:"description,nullzero"`
	Price       float32 `bun:"price,nullzero"`

	OrderItems []*OrderItem `bun:"rel:has-many,join:id=item_id"`
}

type ItemBuilder struct {
	title       string
	description string
	price       float32
}

func NewItemBuilder() *ItemBuilder {
	return &ItemBuilder{}
}

// Title sets the title on the Builder.
func (b *ItemBuilder) Title(title string) *ItemBuilder {
	b.title = title
	return b
}

// Description sets the description on the Builder.
func (b *ItemBuilder) Description(description string) *ItemBuilder {
	b.description = description
	return b
}

// Price sets the price on the Builder.
func (b *ItemBuilder) Price(price float32) *ItemBuilder {
	b.price = price
	return b
}

// Build creates a new Item entity.
func (b *ItemBuilder) Build() *Item {
	return &Item{
		Entity:      Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Title:       b.title,
		Description: b.description,
		Price:       b.price,
	}
}
