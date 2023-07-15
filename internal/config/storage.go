package config

import (
	"github.com/Vaansh/gore/internal/util"
)

type StorageConfig struct {
	BucketName      string
	CredentialsPath string
}

func ReadStorageConfig() *StorageConfig {
	config := &StorageConfig{
		BucketName:      util.Getenv("BUCKET_NAME", true),
		CredentialsPath: util.Getenv("GOOGLE_APPLICATION_CREDENTIALS", true),
	}

	return config
}
