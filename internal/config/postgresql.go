package config

import (
	"cloud.google.com/go/compute/metadata"
	"github.com/Vaansh/gore/internal/util"
)

type DBConfig struct {
	Username   string
	Password   string
	Database   string
	InstanceId string
}

func ReadDbConfig() (*DBConfig, error) {
	instanceId := util.Getenv("INSTANCE_CONNECTION_NAME", false)
	if instanceId == "" {
		var err error
		if instanceId, err = metadata.ProjectID(); err != nil {
			return nil, err
		}
	}

	config := &DBConfig{
		Username:   util.Getenv("DB_USER", true),
		Password:   util.Getenv("DB_PASS", true),
		Database:   util.Getenv("DB_NAME", true),
		InstanceId: instanceId,
	}

	return config, nil
}
