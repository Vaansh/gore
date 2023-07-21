package config

import (
	"strings"

	"github.com/Vaansh/gore/internal/util"
)

type LoggerConfig struct {
	LogName         string
	LocalLog        bool
	CloudLog        bool
	ProjectId       string
	CredentialsPath string
}

func ReadLoggerConfig() *LoggerConfig {
	config := &LoggerConfig{
		LogName:         util.Getenv("LOG_NAME", true),
		LocalLog:        strings.ToLower(util.Getenv("LOCAL_LOG", true)) == "true",
		CloudLog:        strings.ToLower(util.Getenv("CLOUD_LOG", true)) == "true",
		ProjectId:       util.Getenv("PROJECT_ID", true),
		CredentialsPath: util.Getenv("GOOGLE_APPLICATION_CREDENTIALS", true),
	}

	return config
}
