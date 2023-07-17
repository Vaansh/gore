package gcloud

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/Vaansh/gore"
	"github.com/Vaansh/gore/internal/config"
	"github.com/kkdai/youtube/v2"
	"google.golang.org/api/option"
	"io"
	"os"
	"time"
)

var (
	storageClient *storage.Client
	bucketName    string
)

const (
	LocalDataDirectory = "data"
)

func InitStorage() error {
	var err error
	ctx := context.Background()
	cfg := config.ReadStorageConfig()
	bucketName = cfg.BucketName
	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile(cfg.CredentialsPath))
	if err != nil {
		LogFatal(fmt.Sprintf("Failed to create api: %v", err))
	}
	return nil
}

func UploadToBucket(fileName string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	file, err := os.Open(LocalDataDirectory + "/" + fileName)
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}

	wc := storageClient.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %w source: %s", err, "storage bucket")
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %w source: %s", err, "storage bucket")
	}

	return nil
}

func DeleteFromBucket(fileName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := client.Bucket(bucketName).Object(fileName)
	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("Object(%q).Delete: %w", fileName, err)
	}

	return nil
}

func SaveFile(id string, platform gore.Platform) error {
	if platform == gore.YOUTUBE {
		client := youtube.Client{}
		videoLink, err := client.GetVideo(id)
		if err != nil {
			return fmt.Errorf("client.GetVideo(%q): %w", id, err)
		}

		formats := videoLink.Formats.WithAudioChannels()
		stream, _, err := client.GetStream(videoLink, &formats[0])
		if err != nil {
			return fmt.Errorf("client.GetStream(%q): %w", videoLink, err)
		}

		file, err := os.Create(fmt.Sprintf("%s/yt_%s.mp4", LocalDataDirectory, id))
		if err != nil {
			return fmt.Errorf("os.Create(yt_%q.mp4): %w", id, err)
		}

		_, err = io.Copy(file, stream)
		if err != nil {
			LogInfo(fmt.Sprintf("io.Copy: %s source: %s", err, "youtube downloader client"))
		}

		return nil
	} else {
		return fmt.Errorf("platform (%s) file saving not supported", platform)
	}
}

func DeleteFile(fileName string) {
	err := os.Remove(fmt.Sprintf("%s/%s", LocalDataDirectory, fileName))
	if err != nil {
		LogWarning(fmt.Sprintf("Could not delete %s: %s", fileName, err.Error()))
	}
}

func GetFileUrl(fileName string) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, fileName)
}
