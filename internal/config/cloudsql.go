package config

import (
	"cloud.google.com/go/compute/metadata"
	"github.com/Vaansh/gore/internal/util"
	"os"
)

type DBConfig struct {
	Username  string
	Password  string
	Database  string
	ProjectID string
}

func ReadDbConfig() (*DBConfig, error) {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		var err error
		if projectID, err = metadata.ProjectID(); err != nil {
			if err != nil {
				return nil, err
			}
		}
	}

	config := &DBConfig{
		Username:  util.MustGetenv("POSTGRES_USER"),
		Password:  util.MustGetenv("POSTGRES_PASSWORD"),
		Database:  util.MustGetenv("POSTGRES_DB"),
		ProjectID: projectID,
	}

	return config, nil
}
