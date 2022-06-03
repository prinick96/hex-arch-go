package database

import (
	"database/sql"
	"time"

	"github.com/prinick96/elog"
)

type Database interface {
	ConnectDB() *sql.DB
}

// Max seconds for retry a database connection
const DB_CONNECTION_TIMEOUT = 10

// Try db connection
func try(err error, db *sql.DB, counts *int) error {
	if err != nil {
		// increase counter
		elog.New(elog.ERROR, "Trying to connect with database", err)
		*counts++

		// can't connect with the database
		if *counts > DB_CONNECTION_TIMEOUT {
			elog.New(elog.PANIC, "Can't connect with the database", err)
		}

		// log and try again
		elog.New(elog.ERROR, "Backing off for a second", err)
		time.Sleep(time.Second)
	}

	return err
}
