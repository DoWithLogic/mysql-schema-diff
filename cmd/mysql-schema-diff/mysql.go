package main

import (
	"database/sql"
	"fmt"
)

// openDbWithMySQLConfig opens a MySQL database connection using the provided parameters and pings it
func newMySQL(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	if err := conn.Ping(); err != nil {
		conn.Close()

		return nil, err
	}

	return conn, nil
}
