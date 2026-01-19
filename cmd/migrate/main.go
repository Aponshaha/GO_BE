package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"ecom/internal/config"
	"ecom/internal/database"
	"ecom/pkg/migrate"
)

func main() {
	var (
		command = flag.String("command", "up", "Migration command: up, down, status")
		dir     = flag.String("dir", "migrations", "Migrations directory")
	)
	flag.Parse()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Create migration runner
	runner := migrate.NewMigrationRunner(database.GetDB(), *dir)

	// Execute command
	switch *command {
	case "up":
		if err := runner.Up(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	case "down":
		if err := runner.Down(); err != nil {
			log.Fatalf("Rollback failed: %v", err)
		}
	case "status":
		if err := runner.Status(); err != nil {
			log.Fatalf("Status check failed: %v", err)
		}
	default:
		fmt.Println("Usage: go run cmd/migrate/main.go -command=<up|down|status>")
		fmt.Println("\nCommands:")
		fmt.Println("  up     - Run all pending migrations")
		fmt.Println("  down   - Rollback last migration")
		fmt.Println("  status - Show migration status")
		os.Exit(1)
	}
}


