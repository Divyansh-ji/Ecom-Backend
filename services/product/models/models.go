package models

import (
	"time"

	"github.com/google/uuid"
)

// Category represents a product category
type Category struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name" validate:"required,min=1,max=255"`
	Slug        string     `json:"slug" db:"slug" validate:"required,min=1,max=255"`
	Description *string    `json:"description,omitempty" db:"description"`
	ParentID    *uuid.UUID `json:"parent_id,omitempty" db:"parent_id"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

// Product represents a product in the catalog
type Product struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	CategoryID     *uuid.UUID `json:"category_id,omitempty" db:"category_id"`
	Name           string     `json:"name" db:"name" validate:"required,min=1,max=500"`
	Slug           string     `json:"slug" db:"slug" validate:"required,min=1,max=500"`
	Description    *string    `json:"description,omitempty" db:"description"`
	SKU            *string    `json:"sku,omitempty" db:"sku" validate:"omitempty,max=100"`
	Price          float64    `json:"price" db:"price" validate:"required,min=0"`
	CompareAtPrice *float64   `json:"compare_at_price,omitempty" db:"compare_at_price" validate:"omitempty,min=0"`
	Cost           *float64   `json:"cost,omitempty" db:"cost" validate:"omitempty,min=0"`
	Status         string     `json:"status" db:"status" validate:"required,oneof=draft active archived"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
}

// ProductAttribute represents a product attribute (key-value pair)
type ProductAttribute struct {
	ID        uuid.UUID `json:"id" db:"id"`
	ProductID uuid.UUID `json:"product_id" db:"product_id" validate:"required"`
	Key       string    `json:"key" db:"key" validate:"required,min=1,max=100"`
	Value     string    `json:"value" db:"value" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// ProductImage represents a product image
type ProductImage struct {
	ID        uuid.UUID `json:"id" db:"id"`
	ProductID uuid.UUID `json:"product_id" db:"product_id" validate:"required"`
	URL       string    `json:"url" db:"url" validate:"required,url,max=1000"`
	AltText   *string   `json:"alt_text,omitempty" db:"alt_text" validate:"omitempty,max=255"`
	SortOrder int       `json:"sort_order" db:"sort_order"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
