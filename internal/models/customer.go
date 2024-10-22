package models

import (
	"time"

	"github.com/gocql/gocql"
)

type Customer struct {
	ID        gocql.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
