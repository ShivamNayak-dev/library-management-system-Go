package database

import (
	"database/sql"
	"fmt"

	"github.com/ShivamNayak-dev/library-management-system/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}