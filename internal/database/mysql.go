package database

import (
	"context"
	"database/sql"
	"fmt"
	"hex-arch-go/env"

	_ "github.com/go-sql-driver/mysql"
)

// DB Class
type MySQL struct {
	ctx    context.Context
	config env.EnvApp
}

// Ping
func (m *MySQL) ping(err error, db *sql.DB) error {
	if err != nil {
		return err
	}

	// try ping
	err = db.PingContext(m.ctx)
	if err != nil {
		return err
	}

	return nil
}

// Connect
func (m *MySQL) connect() (*sql.DB, error) {
	// try open connection
	db, err := sql.Open(
		m.config.DB_ENGINE,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			m.config.DB_USERNAME,
			m.config.DB_PASSWORD,
			m.config.DB_HOST,
			m.config.DB_PORT,
			m.config.DB_DATABASE,
		))

	// try ping
	if m.ping(err, db) != nil {
		return nil, err
	}

	return db, nil
}

// Get connection
func (m *MySQL) ConnectDB() *sql.DB {
	counts := 0

	for {
		db, err := m.connect()

		if try(err, db, &counts) == nil {
			return db
		}
		continue
	}
}

// Constructor
func NewMySQLDatabase(ctx context.Context, ec env.EnvApp) Database {
	return &MySQL{ctx: ctx, config: ec}
}
