package database

import (
	"database/sql"
	"fmt"
	"github.com/Vaansh/gore/internal/config"
)

func ConnectDb() (*sql.DB, error) {
	cfg, err := config.ReadDbConfig()
	if err != nil {
		return nil, err
	}

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=/cloudsql/%s",
		cfg.Username, cfg.Password, cfg.Database, cfg.ProjectID)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
