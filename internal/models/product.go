package models

import (
	"time"

	"github.com/gocql/gocql"
)

type Product struct {
	ID          gocql.UUID
	Name        string
	Description string
	Price       uint32
	Stock       uint32
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
