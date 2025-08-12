package di

import (
	"articles/shared/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type (
	PostgresdbOption struct {
		ConnectionString                     string
		MaxLifeTimeConnection                time.Duration
		MaxIdleConnection, MaxOpenConnection int
	}
)

func NewDB(config *config.Configuration) (*sql.DB, error) {
	// Parse the connection lifetime duration
	duration, err := time.ParseDuration(config.DbMaxLifeTimeConnection)
	if err != nil {
		return nil, err
	}

	// Create the connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPass,
		config.DbName,
	)

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	db.SetMaxOpenConns(config.DbMaxOpenConnection)
	db.SetMaxIdleConns(config.DbMaxIdleConnection)
	db.SetConnMaxLifetime(duration)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
