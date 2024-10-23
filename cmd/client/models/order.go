package models

import (
	"time"

	"github.com/gocql/gocql"
)

type Order struct {
	ID         gocql.UUID `json:"order_id"`
	ProductID  gocql.UUID `json:"product_id"`
	Quantity   uint32     `json:"quantity"`
	CustomerID gocql.UUID `json:"customer_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
