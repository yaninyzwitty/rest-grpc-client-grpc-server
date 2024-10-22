// using astra db

package database

import (
	"log/slog"
	"time"

	gocqlastra "github.com/datastax/gocql-astra"
	"github.com/gocql/gocql"
)

type DatabaseConfig struct {
	Path     string
	Username string
	Password string
	Timeout  time.Duration
}

func NewDatabaseConnection(config DatabaseConfig) (*gocql.Session, error) {
	cluster, err := gocqlastra.NewClusterFromBundle(config.Path, config.Username, config.Password, config.Timeout)
	cluster.Timeout = config.Timeout
	if err != nil {
		slog.Error("unable to load bundle: ", "error", err)
		return nil, err
	}
	session, err := gocql.NewSession(*cluster)
	if err != nil {
		slog.Error("failed to create cluster: ", "error", err)
		return nil, err
	}
	return session, nil
}
