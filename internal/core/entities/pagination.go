package entities

import "fmt"

// Direction can be ASC, DESC, ASC_NULLS_FIRST, DESC_NULLS_FIRST, ASC_NULLS_LAST or DESC_NULLS_LAST.
type Direction string

const (
	ASC              Direction = "ASC"
	DESC             Direction = "DESC"
	ASC_NULLS_FIRST  Direction = "ASC NULLS FIRST"
	DESC_NULLS_FIRST Direction = "DESC NULLS FIRST"
	ASC_NULLS_LAST   Direction = "ASC NULLS LAST"
	DESC_NULLS_LAST  Direction = "DESC NULLS LAST"
)

// SortOrder represent single sort instruction.
type SortOrder struct {
	Property  string
	Direction Direction
}

// OrderOption function is used for creating new SortOrder object.
type OrderOption func(*SortOrder)

// NewSortOrder creates new SortOrder object.
func NewSortOrder(opts ...OrderOption) *SortOrder {
	order := &SortOrder{Property: "created_at", Direction: DESC}
	for _, opt := range opts {
		opt(order)
	}
	return order
}

func WithDirection(d Direction) OrderOption {
	return func(order *SortOrder) {
		order.Direction = d
	}
}

func WithProperty(p string) OrderOption {
	return func(order *SortOrder) {
		order.Property = p
	}
}

// Sort represents sorting options for pagination.
type Sort struct {
	Orders []*SortOrder
}

// NewSort creates new sort object from Orders.
func NewSort(order ...*SortOrder) Sort {
	return Sort{Orders: order}
}

// Pageable represents the pagination request parameters.
type Pageable struct {
	Size   int
	Offset int
	Sort   Sort
}

// StringifyOrders travers sort orders into a slice of strings in the following of "prop1 ASC, prop2 DESC".
func StringifyOrders(s Sort) []string {
	var orders []string
	for _, o := range s.Orders {
		orders = append(orders, fmt.Sprintf("%s %s", o.Property, o.Direction))
	}
	return orders
}
