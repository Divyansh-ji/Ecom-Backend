package models

import (
	"time"

	"github.com/google/uuid"
)

// Warehouse represents a warehouse location
type Warehouse struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required,min=1,max=255"`
	Code      string    `json:"code" db:"code" validate:"required,min=1,max=50"`
	Address   *string   `json:"address,omitempty" db:"address"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Stock represents inventory stock for a product in a warehouse
type Stock struct {
	ID          uuid.UUID `json:"id" db:"id"`
	ProductID   uuid.UUID `json:"product_id" db:"product_id" validate:"required"`
	WarehouseID uuid.UUID `json:"warehouse_id" db:"warehouse_id" validate:"required"`
	Quantity    int       `json:"quantity" db:"quantity" validate:"min=0"`
	Reserved    int       `json:"reserved" db:"reserved" validate:"min=0"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// StockMovementType represents the type of stock movement
type StockMovementType string

const (
	StockMovementIn       StockMovementType = "in"
	StockMovementOut      StockMovementType = "out"
	StockMovementAdjust   StockMovementType = "adjust"
	StockMovementTransfer StockMovementType = "transfer"
	StockMovementReserve  StockMovementType = "reserve"
	StockMovementRelease  StockMovementType = "release"
)

// StockMovement represents a stock movement transaction
type StockMovement struct {
	ID            uuid.UUID         `json:"id" db:"id"`
	StockID       uuid.UUID         `json:"stock_id" db:"stock_id" validate:"required"`
	Type          StockMovementType `json:"type" db:"type" validate:"required,oneof=in out adjust transfer reserve release"`
	Quantity      int               `json:"quantity" db:"quantity" validate:"required"`
	ReferenceID   *uuid.UUID        `json:"reference_id,omitempty" db:"reference_id"`
	ReferenceType *string           `json:"reference_type,omitempty" db:"reference_type" validate:"omitempty,max=50"`
	Reason        *string           `json:"reason,omitempty" db:"reason" validate:"omitempty,max=255"`
	CreatedAt     time.Time         `json:"created_at" db:"created_at"`
}

// StockReservationStatus represents the status of a stock reservation
type StockReservationStatus string

const (
	ReservationStatusActive    StockReservationStatus = "active"
	ReservationStatusFulfilled StockReservationStatus = "fulfilled"
	ReservationStatusExpired   StockReservationStatus = "expired"
	ReservationStatusCancelled StockReservationStatus = "cancelled"
)

// StockReservation represents a reserved stock for an order
type StockReservation struct {
	ID         uuid.UUID              `json:"id" db:"id"`
	StockID    uuid.UUID              `json:"stock_id" db:"stock_id" validate:"required"`
	OrderID    uuid.UUID              `json:"order_id" db:"order_id" validate:"required"`
	Quantity   int                    `json:"quantity" db:"quantity" validate:"required,min=1"`
	ExpiresAt  time.Time              `json:"expires_at" db:"expires_at" validate:"required"`
	Status     StockReservationStatus `json:"status" db:"status" validate:"required,oneof=active fulfilled expired cancelled"`
	CreatedAt  time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at" db:"updated_at"`
}
