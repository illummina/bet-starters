package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var db *sqlx.DB

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func Init() (*sqlx.DB, error) {
	if db != nil {
		return db, nil
	}

	cfg := &Config{
		os.Getenv("MSSQL_CONTAINER_NAME"),
		os.Getenv("MSSQL_PORT"),
		os.Getenv("MSSQL_USER"),
		os.Getenv("MSSQL_PASSWORD"),
		os.Getenv("MSSQL_DATABASE"),
	}

	fmt.Println(cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	if cfg.Host == "" || cfg.Port == "" || cfg.User == "" || cfg.Password == "" || cfg.Database == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?db=%s&connection+timeout=30", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	var err error
	db, err = sqlx.Connect("sqlserver", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %v", err)
	}

	log.Println("Database initialized successfully!")
	return db, nil
}
