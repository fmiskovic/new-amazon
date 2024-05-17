package entities

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

// OrderItem handles the many-to-many relationship between orders and items,
// recording which items are in which orders and in what quantities.
type OrderItem struct {
	bun.BaseModel `bun:"table:order_items,alias:oi"`

	Entity

	OrderID uuid.UUID `bun:"order_id,notnull"`
	Order   *Order    `bun:"rel:belongs-to,join:order_id=id"`

	ItemID uuid.UUID `bun:"item_id,notnull"`
	Item   *Item     `bun:"rel:belongs-to,join:item_id=id"`

	Quantity int `bun:"quantity,notnull"`
}

type OrderItemBuilder struct {
	orderID  uuid.UUID
	itemID   uuid.UUID
	quantity int
}

func NewOrderItemBuilder() *OrderItemBuilder {
	return &OrderItemBuilder{}
}

func (b *OrderItemBuilder) OrderID(orderID uuid.UUID) *OrderItemBuilder {
	b.orderID = orderID
	return b
}

func (b *OrderItemBuilder) ItemID(itemID uuid.UUID) *OrderItemBuilder {
	b.itemID = itemID
	return b
}

func (b *OrderItemBuilder) Quantity(quantity int) *OrderItemBuilder {
	b.quantity = quantity
	return b
}

func (b *OrderItemBuilder) Build() *OrderItem {
	return &OrderItem{
		Entity:   Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		OrderID:  b.orderID,
		ItemID:   b.itemID,
		Quantity: b.quantity,
	}
}
