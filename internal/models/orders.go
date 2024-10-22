package models

import (
	"time"

	"github.com/gocql/gocql"
)

type Order struct {
	ID         gocql.UUID
	ProductID  gocql.UUID
	Quantity   uint32
	CustomerID gocql.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
