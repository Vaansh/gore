package config

import (
	"cloud.google.com/go/compute/metadata"
	"database/sql"
	"fmt"
	"github.com/Vaansh/gore/internal/util"
	"os"
)

func ConnectToDB() (*sql.DB, error) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		var err error
		if projectID, err = metadata.ProjectID(); err != nil {
			if err != nil {
				return nil, err
			}
		}
	}

	username := util.MustGetenv("POSTGRES_USER")
	password := util.MustGetenv("POSTGRES_PASSWORD")
	dbName := util.MustGetenv("POSTGRES_DB")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=/cloudsql/%s", username, password, dbName, projectID)

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
