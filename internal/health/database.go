package health

import (
	"context"
	"database/sql"
)

type DatabaseChecker struct {
	db *sql.DB
}

func NewDatabaseChecker(db *sql.DB) *DatabaseChecker {
	return &DatabaseChecker{db: db}
}

func (d *DatabaseChecker) Check(ctx context.Context) error {
	// Example: Ping the database.
	return d.db.PingContext(ctx)
}
