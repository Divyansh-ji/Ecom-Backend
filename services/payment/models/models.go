package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// PaymentMethodType represents the type of payment method
type PaymentMethodType string

const (
	PaymentMethodCard   PaymentMethodType = "card"
	PaymentMethodBank   PaymentMethodType = "bank"
	PaymentMethodWallet PaymentMethodType = "wallet"
	PaymentMethodCOD    PaymentMethodType = "cod"
	PaymentMethodOther  PaymentMethodType = "other"
)

// PaymentMethod represents a user's payment method
type PaymentMethod struct {
	ID         uuid.UUID         `json:"id" db:"id"`
	UserID     uuid.UUID         `json:"user_id" db:"user_id" validate:"required"`
	Type       PaymentMethodType `json:"type" db:"type" validate:"required,oneof=card bank wallet cod other"`
	Provider   *string           `json:"provider,omitempty" db:"provider" validate:"omitempty,max=50"`
	LastFour   *string           `json:"last_four,omitempty" db:"last_four" validate:"omitempty,len=4"`
	Expiry     *string           `json:"expiry,omitempty" db:"expiry" validate:"omitempty,len=7"`
	IsDefault  bool              `json:"is_default" db:"is_default"`
	Metadata   json.RawMessage   `json:"metadata,omitempty" db:"metadata"`
	CreatedAt  time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at" db:"updated_at"`
}

// PaymentStatus represents the status of a payment
type PaymentStatus string

const (
	PaymentStatusPending           PaymentStatus = "pending"
	PaymentStatusAuthorized        PaymentStatus = "authorized"
	PaymentStatusCaptured          PaymentStatus = "captured"
	PaymentStatusFailed            PaymentStatus = "failed"
	PaymentStatusRefunded          PaymentStatus = "refunded"
	PaymentStatusPartiallyRefunded PaymentStatus = "partially_refunded"
	PaymentStatusCancelled         PaymentStatus = "cancelled"
)

// Payment represents a payment transaction
type Payment struct {
	ID              uuid.UUID       `json:"id" db:"id"`
	OrderID         uuid.UUID       `json:"order_id" db:"order_id" validate:"required"`
	UserID          uuid.UUID       `json:"user_id" db:"user_id" validate:"required"`
	Amount          float64         `json:"amount" db:"amount" validate:"required,min=0.01"`
	Currency        string          `json:"currency" db:"currency" validate:"required,len=3"`
	Status          PaymentStatus   `json:"status" db:"status" validate:"required,oneof=pending authorized captured failed refunded partially_refunded cancelled"`
	PaymentMethodID *uuid.UUID      `json:"payment_method_id,omitempty" db:"payment_method_id"`
	Gateway         *string         `json:"gateway,omitempty" db:"gateway" validate:"omitempty,max=50"`
	GatewayRef      *string         `json:"gateway_ref,omitempty" db:"gateway_ref" validate:"omitempty,max=255"`
	GatewayResponse json.RawMessage `json:"gateway_response,omitempty" db:"gateway_response"`
	CreatedAt       time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
}

// PaymentTransactionType represents the type of payment transaction
type PaymentTransactionType string

const (
	TransactionTypeAuth       PaymentTransactionType = "auth"
	TransactionTypeCapture    PaymentTransactionType = "capture"
	TransactionTypeRefund     PaymentTransactionType = "refund"
	TransactionTypeVoid       PaymentTransactionType = "void"
	TransactionTypeAdjustment PaymentTransactionType = "adjustment"
)

// TransactionStatus represents the status of a transaction
type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "pending"
	TransactionStatusSuccess TransactionStatus = "success"
	TransactionStatusFailed  TransactionStatus = "failed"
)

// PaymentTransaction represents a payment transaction record
type PaymentTransaction struct {
	ID              uuid.UUID              `json:"id" db:"id"`
	PaymentID       uuid.UUID              `json:"payment_id" db:"payment_id" validate:"required"`
	Type            PaymentTransactionType `json:"type" db:"type" validate:"required,oneof=auth capture refund void adjustment"`
	Amount          float64                `json:"amount" db:"amount" validate:"required"`
	Status          TransactionStatus       `json:"status" db:"status" validate:"required,oneof=pending success failed"`
	GatewayTxnID    *string                `json:"gateway_txn_id,omitempty" db:"gateway_txn_id" validate:"omitempty,max=255"`
	GatewayResponse json.RawMessage        `json:"gateway_response,omitempty" db:"gateway_response"`
	CreatedAt       time.Time              `json:"created_at" db:"created_at"`
}

// RefundStatus represents the status of a refund
type RefundStatus string

const (
	RefundStatusPending   RefundStatus = "pending"
	RefundStatusCompleted RefundStatus = "completed"
	RefundStatusFailed    RefundStatus = "failed"
)

// Refund represents a refund transaction
type Refund struct {
	ID        uuid.UUID   `json:"id" db:"id"`
	PaymentID uuid.UUID   `json:"payment_id" db:"payment_id" validate:"required"`
	Amount    float64     `json:"amount" db:"amount" validate:"required,min=0.01"`
	Reason    *string     `json:"reason,omitempty" db:"reason" validate:"omitempty,max=255"`
	Status    RefundStatus `json:"status" db:"status" validate:"required,oneof=pending completed failed"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
}
