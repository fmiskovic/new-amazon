package entities

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

// Order will store information about each order made by an account.
type Order struct {
	bun.BaseModel `bun:"table:orders,alias:o"`

	Entity

	// many-to-one relation
	AccountID uuid.UUID `bun:"account_id,notnull"`
	Account   Account   `bun:"rel:belongs-to,join:account_id=id"`

	OrderItems []*OrderItem `bun:"rel:has-many,join:id=order_id"`
}

// OrderBuilder is a builder pattern for creating new Order entities.
type OrderBuilder struct {
	accountID  uuid.UUID
	account    Account
	orderItems []*OrderItem
}

// NewOrderBuilder creates a new OrderBuilder.
func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{}
}

// AccountID sets the accountID on the Builder.
func (b *OrderBuilder) AccountID(accountID uuid.UUID) *OrderBuilder {
	b.accountID = accountID
	return b
}

// OrderItems sets the orderItems on the Builder.
func (b *OrderBuilder) OrderItems(orderItems []*OrderItem) *OrderBuilder {
	b.orderItems = orderItems
	return b
}

// Build creates a new Order entity.
func (b *OrderBuilder) Build() *Order {
	return &Order{
		Entity:     Entity{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
		AccountID:  b.accountID,
		Account:    b.account,
		OrderItems: b.orderItems,
	}
}
