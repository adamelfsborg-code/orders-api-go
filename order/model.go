package order

import (
	"time"

	"github.com/google/uuid"
)

type OrderModel struct {
	OrderID     uint64               `json:"order_id"`
	CustomerID  uuid.UUID            `json:"customer_id"`
	LineItems   []OrderLineItemModel `json:"line_items"`
	CreatedAt   *time.Time           `json:"created_at"`
	ShippedAt   *time.Time           `json:"shipped_at"`
	CompletedAt *time.Time           `json:"completed_at"`
}

type OrderLineItemModel struct {
	ItemID   uuid.UUID `json:"item_id"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
