package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusConfirmed  OrderStatus = "confirmed"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusRefunded   OrderStatus = "refunded"
)

// Address represents a shipping or billing address
type Address struct {
	FirstName  string  `json:"first_name" validate:"required"`
	LastName   string  `json:"last_name" validate:"required"`
	Company    *string `json:"company,omitempty"`
	Address1   string  `json:"address1" validate:"required"`
	Address2   *string `json:"address2,omitempty"`
	City       string  `json:"city" validate:"required"`
	State      string  `json:"state" validate:"required"`
	PostalCode string  `json:"postal_code" validate:"required"`
	Country    string  `json:"country" validate:"required"`
	Phone      *string `json:"phone,omitempty"`
	Email      *string `json:"email,omitempty" validate:"omitempty,email"`
}

// Order represents an order
type Order struct {
	ID              uuid.UUID       `json:"id" db:"id"`
	UserID          uuid.UUID       `json:"user_id" db:"user_id" validate:"required"`
	OrderNumber     string          `json:"order_number" db:"order_number" validate:"required,min=1,max=50"`
	Status          OrderStatus     `json:"status" db:"status" validate:"required,oneof=pending confirmed processing shipped delivered cancelled refunded"`
	Subtotal        float64         `json:"subtotal" db:"subtotal" validate:"required,min=0"`
	TaxAmount       float64         `json:"tax_amount" db:"tax_amount" validate:"min=0"`
	ShippingAmount  float64         `json:"shipping_amount" db:"shipping_amount" validate:"min=0"`
	DiscountAmount  float64         `json:"discount_amount" db:"discount_amount" validate:"min=0"`
	Total           float64         `json:"total" db:"total" validate:"required,min=0"`
	Currency        string          `json:"currency" db:"currency" validate:"required,len=3"`
	ShippingAddress json.RawMessage `json:"shipping_address" db:"shipping_address"`
	BillingAddress  json.RawMessage `json:"billing_address" db:"billing_address"`
	Notes           *string         `json:"notes,omitempty" db:"notes"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID         uuid.UUID       `json:"id" db:"id"`
	OrderID    uuid.UUID       `json:"order_id" db:"order_id" validate:"required"`
	ProductID  uuid.UUID       `json:"product_id" db:"product_id" validate:"required"`
	SKU        *string         `json:"sku,omitempty" db:"sku" validate:"omitempty,max=100"`
	Name       string          `json:"name" db:"name" validate:"required,min=1,max=500"`
	Quantity   int             `json:"quantity" db:"quantity" validate:"required,min=1"`
	UnitPrice  float64         `json:"unit_price" db:"unit_price" validate:"required,min=0"`
	TotalPrice float64         `json:"total_price" db:"total_price" validate:"required,min=0"`
	Metadata   json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	CreatedAt  time.Time       `json:"created_at" db:"created_at"`
}

// OrderStatusHistory represents a status change in an order
type OrderStatusHistory struct {
	ID        uuid.UUID   `json:"id" db:"id"`
	OrderID   uuid.UUID   `json:"order_id" db:"order_id" validate:"required"`
	Status    OrderStatus `json:"status" db:"status" validate:"required"`
	Note      *string     `json:"note,omitempty" db:"note"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
}

// OrderShipment represents shipment information for an order
type OrderShipment struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	OrderID     uuid.UUID  `json:"order_id" db:"order_id" validate:"required"`
	Carrier     *string    `json:"carrier,omitempty" db:"carrier" validate:"omitempty,max=100"`
	TrackingNo  *string    `json:"tracking_no,omitempty" db:"tracking_no" validate:"omitempty,max=255"`
	ShippedAt   *time.Time `json:"shipped_at,omitempty" db:"shipped_at"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty" db:"delivered_at"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}
