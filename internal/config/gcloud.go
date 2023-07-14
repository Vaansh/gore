package config

import (
	"github.com/Vaansh/gore/internal/util"
)

type BucketConfig struct {
	BucketName      string
	CredentialsPath string
}

func ReadBucketConfig() *BucketConfig {
	config := &BucketConfig{
		BucketName:      util.Getenv("BUCKET_NAME", true),
		CredentialsPath: util.Getenv("GOOGLE_APPLICATION_CREDENTIALS", true),
	}

	return config
}
