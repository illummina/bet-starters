package db

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"path/filepath"
)

func Migrate(db *sqlx.DB) error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
	}

	sourceURL := filepath.Join(filepath.Dir(filepath.Dir(wd)), "internal", "db", "migrations")

	driver, err := sqlserver.WithInstance(db.DB, &sqlserver.Config{})
	if err != nil {
		return fmt.Errorf("failed to create sqlserver driver instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+sourceURL+"/", "sqlserver", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %v", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}

	log.Println("Migrations completed successfully!")
	return nil
}
