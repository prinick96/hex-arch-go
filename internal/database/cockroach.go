package database

import (
	"context"
	"database/sql"
	"fmt"
	"hex-arch-go/env"

	_ "github.com/lib/pq"
)

// DB Class
type CockRoach struct {
	ctx    context.Context
	config env.EnvApp
}

// Ping
func (c *CockRoach) ping(err error, db *sql.DB) error {
	if err != nil {
		return err
	}

	// try ping
	err = db.PingContext(c.ctx)
	if err != nil {
		return err
	}

	return nil
}

// Connect
func (c *CockRoach) connect() (*sql.DB, error) {
	// try open connection
	db, err := sql.Open(
		c.config.DB_ENGINE,
		fmt.Sprintf(
			"%s://%s:%s@%s:%s/%s?sslmode=verify-full&options=%s",
			c.config.DB_ENGINE,
			c.config.DB_USERNAME,
			c.config.DB_PASSWORD,
			c.config.DB_HOST,
			c.config.DB_PORT,
			c.config.DB_DATABASE,
			c.config.DB_OPTIONS,
		))

	// try ping
	if c.ping(err, db) != nil {
		return nil, err
	}

	return db, nil
}

// Get connection
func (c *CockRoach) ConnectDB() *sql.DB {
	counts := 0

	for {
		db, err := c.connect()

		if try(err, db, &counts) == nil {
			return db
		}
		continue
	}
}

// Constructor
func NewCockRoachDatabase(ctx context.Context, ec env.EnvApp) Database {
	return &CockRoach{ctx: ctx, config: ec}
}
