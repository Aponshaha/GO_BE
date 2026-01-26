package migrate

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/lib/pq"
)

// MigrationRunner handles database migrations
type MigrationRunner struct {
	db    *sql.DB
	dir   string
	table string
}

// NewMigrationRunner creates a new migration runner
func NewMigrationRunner(db *sql.DB, migrationsDir string) *MigrationRunner {
	return &MigrationRunner{
		db:    db,
		dir:   migrationsDir,
		table: "schema_migrations",
	}
}

// Init creates the migrations tracking table
func (m *MigrationRunner) Init() error {
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`, m.table)

	_, err := m.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	log.Println("✅ Migrations table initialized")
	return nil
}

// GetAppliedMigrations returns a map of applied migration versions
func (m *MigrationRunner) GetAppliedMigrations() (map[string]bool, error) {
	applied := make(map[string]bool)

	query := fmt.Sprintf("SELECT version FROM %s ORDER BY version", m.table)
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query applied migrations: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		applied[version] = true
	}

	return applied, nil
}

// GetMigrationFiles returns sorted list of migration files
func (m *MigrationRunner) GetMigrationFiles() ([]string, error) {
	files, err := os.ReadDir(m.dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read migrations directory: %w", err)
	}

	var migrations []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasSuffix(file.Name(), ".sql") {
			migrations = append(migrations, file.Name())
		}
	}

	// Sort by filename (which should be numbered)
	sort.Strings(migrations)
	return migrations, nil
}

// RunMigration executes a single migration file
func (m *MigrationRunner) RunMigration(filename string) error {
	filepath := filepath.Join(m.dir, filename)

	// Read migration file
	content, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read migration file %s: %w", filename, err)
	}

	// Extract version from filename (e.g., "001_create_tables.sql" -> "001")
	version := strings.Split(filename, "_")[0]

	// Start transaction
	tx, err := m.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Execute migration
	if _, err := tx.Exec(string(content)); err != nil {
		return fmt.Errorf("failed to execute migration %s: %w", filename, err)
	}

	// Record migration
	insertQuery := fmt.Sprintf(
		"INSERT INTO %s (version) VALUES ($1) ON CONFLICT (version) DO NOTHING",
		m.table,
	)
	if _, err := tx.Exec(insertQuery, version); err != nil {
		return fmt.Errorf("failed to record migration %s: %w", version, err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit migration %s: %w", version, err)
	}

	log.Printf("✅ Applied migration: %s", filename)
	return nil
}

// Up runs all pending migrations
func (m *MigrationRunner) Up() error {
	// Initialize migrations table
	if err := m.Init(); err != nil {
		return err
	}

	// Get applied migrations
	applied, err := m.GetAppliedMigrations()
	if err != nil {
		return err
	}

	// Get all migration files
	files, err := m.GetMigrationFiles()
	if err != nil {
		return err
	}

	// Run pending migrations
	for _, file := range files {
		version := strings.Split(file, "_")[0]

		if applied[version] {
			log.Printf("⏭️  Skipping already applied migration: %s", file)
			continue
		}

		if err := m.RunMigration(file); err != nil {
			return fmt.Errorf("migration failed at %s: %w", file, err)
		}
	}

	log.Println("✅ All migrations applied successfully")
	return nil
}

// Status shows migration status
func (m *MigrationRunner) Status() error {
	// Initialize if needed
	if err := m.Init(); err != nil {
		return err
	}

	applied, err := m.GetAppliedMigrations()
	if err != nil {
		return err
	}

	files, err := m.GetMigrationFiles()
	if err != nil {
		return err
	}

	fmt.Println("\n📊 Migration Status:")
	fmt.Println(strings.Repeat("-", 60))

	for _, file := range files {
		version := strings.Split(file, "_")[0]
		if applied[version] {
			fmt.Printf("✅ %s (applied)\n", file)
		} else {
			fmt.Printf("⏳ %s (pending)\n", file)
		}
	}
	fmt.Println(strings.Repeat("-", 60))

	return nil
}

// Down rolls back the last migration (simple implementation)
// Note: For production, you'd want proper rollback migrations
func (m *MigrationRunner) Down() error {
	applied, err := m.GetAppliedMigrations()
	if err != nil {
		return err
	}

	if len(applied) == 0 {
		log.Println("ℹ️  No migrations to rollback")
		return nil
	}

	// Get all files
	files, err := m.GetMigrationFiles()
	if err != nil {
		return err
	}

	// Find last applied migration
	var lastFile string
	for i := len(files) - 1; i >= 0; i-- {
		version := strings.Split(files[i], "_")[0]
		if applied[version] {
			lastFile = files[i]
			break
		}
	}

	if lastFile == "" {
		log.Println("ℹ️  No migrations to rollback")
		return nil
	}

	version := strings.Split(lastFile, "_")[0]

	// Remove from tracking (actual rollback would require rollback SQL files)
	query := fmt.Sprintf("DELETE FROM %s WHERE version = $1", m.table)
	if _, err := m.db.Exec(query, version); err != nil {
		return fmt.Errorf("failed to remove migration record: %w", err)
	}

	log.Printf("⚠️  Removed migration record: %s", lastFile)
	log.Println("⚠️  Note: This only removes the record. Manual rollback may be needed.")

	return nil
}
