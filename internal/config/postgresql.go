package config

import (
	"github.com/Vaansh/gore/internal/util"
)

type DBConfig struct {
	Username   string
	Password   string
	Database   string
	InstanceId string
}

func ReadDbConfig() *DBConfig {
	config := &DBConfig{
		Username:   util.Getenv("DB_USER", true),
		Password:   util.Getenv("DB_PASS", true),
		Database:   util.Getenv("DB_NAME", true),
		InstanceId: util.Getenv("INSTANCE_CONNECTION_NAME", true),
	}

	return config
}
