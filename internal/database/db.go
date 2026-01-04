package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"ecom/internal/config"
)

var DB *sql.DB

// Connect initializes the database connection
func Connect(cfg *config.Config) error {
	var err error

	dsn := cfg.GetDSN()
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("✅ Database connected successfully")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return DB
}

