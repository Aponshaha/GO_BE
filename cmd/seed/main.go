package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"ecom/internal/config"
	"ecom/internal/database"
)

func main() {
	var (
		clear = flag.Bool("clear", false, "Clear all data before seeding")
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

	db := database.GetDB()

	if *clear {
		fmt.Println("🗑️  Clearing existing data...")
		if err := clearData(db); err != nil {
			log.Fatalf("Failed to clear data: %v", err)
		}
		fmt.Println("✅ Data cleared")
	}

	fmt.Println("🌱 Seeding database with mock data...")

	// Run all seed files
	if err := runSeedFiles(db); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	fmt.Println("✅ Database seeding completed successfully!")
}

func clearData(db *sql.DB) error {
	tables := []string{
		"order_items",
		"order_coupons",
		"payments",
		"orders",
		"wishlist_items",
		"wishlists",
		"product_reviews",
		"shipping_details",
		"inventory_movements",
		"product_images",
		"customer_addresses",
		"products",
		"customers",
		"coupons",
		"categories",
	}

	for _, table := range tables {
		if _, err := db.Exec(fmt.Sprintf("DELETE FROM %s CASCADE", table)); err != nil {
			// Continue even if some tables don't exist
			continue
		}
	}

	return nil
}

func runSeedFiles(db *sql.DB) error {
	// Get seed files from seeds directory
	seedDir := "seeds"
	files, err := os.ReadDir(seedDir)
	if err != nil {
		return fmt.Errorf("failed to read seeds directory: %w", err)
	}

	// Filter and sort SQL files
	var sqlFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}
	sort.Strings(sqlFiles)

	if len(sqlFiles) == 0 {
		fmt.Println("  ⚠️  No seed files found in seeds directory")
		return nil
	}

	// Execute each seed file
	for _, filename := range sqlFiles {
		filepath := filepath.Join(seedDir, filename)
		fmt.Printf("  📄 Running %s...\n", filename)

		content, err := os.ReadFile(filepath)
		if err != nil {
			return fmt.Errorf("failed to read seed file %s: %w", filename, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to execute seed file %s: %w", filename, err)
		}

		fmt.Printf("    ✓ %s completed\n", filename)
	}

	return nil
}
