package util

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"log"
	"os"
)

const (
	DIRECTORY = "data"
)

func MustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Fatal Error: %s environment variable not set.\n", k)
	}
	return v
}

func SaveYoutubeVideo(id string) {
	http := youtube.Client{}

	videoLink, err := http.GetVideo(id)
	if err != nil {
	}

	formats := videoLink.Formats.WithAudioChannels()
	stream, _, err := http.GetStream(videoLink, &formats[0])
	if err != nil {
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.mp4", DIRECTORY, id))
	if err != nil {
	}

	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
	}
}

func Delete(id string) {
	err := os.Remove(fmt.Sprintf("%s/%s.mp4", DIRECTORY, id))

	if err != nil {
	}

	fmt.Println("File deleted successfully!")
}
