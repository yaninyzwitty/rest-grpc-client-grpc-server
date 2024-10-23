package models

import (
	"time"

	"github.com/gocql/gocql"
)

type Product struct {
	ID          gocql.UUID `json:"product_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       uint32     `json:"price"`
	Stock       uint32     `json:"stock"`
	Category    string     `json:"category"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
