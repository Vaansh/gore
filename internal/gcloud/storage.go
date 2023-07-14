package gcloud

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/Vaansh/gore/internal/config"
	"github.com/kkdai/youtube/v2"
	"google.golang.org/api/option"
	"io"
	"log"
	"os"
)

const (
	DIRECTORY = "data"
)

type StorageHandler struct {
	storageClient *storage.Client
	youtubeClient youtube.Client
	bucketName    string
}

func NewStorageHandler() *StorageHandler {
	ctx := context.Background()
	cfg := config.ReadBucketConfig()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(cfg.CredentialsPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &StorageHandler{
		storageClient: client,
		youtubeClient: youtube.Client{},
		bucketName:    cfg.BucketName,
	}
}

func (fh *StorageHandler) UploadToBucket(fileName string) bool {
	ctx := context.Background()
	file, err := os.Open(fmt.Sprintf("%s/%s", DIRECTORY, fileName))
	if err != nil {
		//Log
		log.Println(err)
		return false
	}
	defer file.Close()

	wc := fh.storageClient.Bucket(fh.bucketName).Object(fileName).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		log.Println(err)
		return false
	}

	if err := wc.Close(); err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (fh *StorageHandler) SaveYoutubeVideo(id string) bool {
	videoLink, err := fh.youtubeClient.GetVideo(id)
	if err != nil {
		//Log
		return false
	}

	formats := videoLink.Formats.WithAudioChannels()
	stream, _, err := fh.youtubeClient.GetStream(videoLink, &formats[0])
	if err != nil {
		return false
	}

	file, err := os.Create(fmt.Sprintf("%s/yt_%s.mp4", DIRECTORY, id))
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return false
	}
	return true
}

func (fh *StorageHandler) GetFileUrl(fileName string) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", fh.bucketName, fileName)
}

func Delete(id string) {
	err := os.Remove(fmt.Sprintf("%s/%s.mp4", DIRECTORY, id))
	if err != nil {
	}

	fmt.Println("File deleted successfully!")
}
