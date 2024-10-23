package models

import (
	"time"

	"github.com/gocql/gocql"
)

type Customer struct {
	ID        gocql.UUID `json:"customer_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
