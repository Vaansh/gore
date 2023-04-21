package publisher

import (
	"fmt"
	"github.com/Vaansh/gore/internal/http"
	"time"
)

type YoutubePublisher struct {
	ChannelID string
	client    *http.YoutubeClient
}

func NewPublisher(channelID string) *YoutubePublisher {
	return &YoutubePublisher{
		ChannelID: channelID,
		client:    http.NewYoutubeClient("ApiKey"),
	}
}

func (p *YoutubePublisher) PublishTo(c chan<- string) {
	firstURL := p.client.PaginatedVideosAPI(p.ChannelID)
	fmt.Println(firstURL)

	videoURL := firstURL.String()
	for {
		data, err := p.client.FetchVideo(videoURL)
		if err != nil {
			fmt.Println(err)
			break
		}

		for _, item := range data["items"].([]interface{}) {
			itemMap := item.(map[string]interface{})
			if itemMap["id"].(map[string]interface{})["kind"] == "youtube#video" {
				videoID := itemMap["id"].(map[string]interface{})["videoId"].(string)
				c <- videoID
			}
		}

		if nextPageToken, ok := data["nextPageToken"].(string); ok {
			videoURL = fmt.Sprintf("%s&pageToken=%s", firstURL, nextPageToken)
		} else {
			break
		}

		time.Sleep(3 * time.Hour)
	}

	var videosBuffer []string
	for {
		mostRecentUpload, err := p.client.FetchLatestVideoByChannel(p.ChannelID)

		if err != nil {
		}

		if !contains(videosBuffer, mostRecentUpload) {
			c <- mostRecentUpload
			videosBuffer = append(videosBuffer, mostRecentUpload)
		}

		if len(videosBuffer) == 50 {
			videosBuffer = make([]string, 0)
		}

		time.Sleep(3 * time.Hour)
	}
}

func (p *YoutubePublisher) PublishVideoTo(c chan<- string) {
	firstURL := p.client.PaginatedVideosAPI(p.ChannelID)
	fmt.Println(firstURL)

	videoURL := firstURL.String()
	for {
		data, err := p.client.FetchVideo(videoURL)
		if err != nil {
			fmt.Println(err)
			break
		}

		for _, item := range data["items"].([]interface{}) {
			itemMap := item.(map[string]interface{})
			if itemMap["id"].(map[string]interface{})["kind"] == "youtube#video" {
				videoID := itemMap["id"].(map[string]interface{})["videoId"].(string)
				c <- videoID
			}
		}

		if nextPageToken, ok := data["nextPageToken"].(string); ok {
			videoURL = fmt.Sprintf("%s&pageToken=%s", firstURL, nextPageToken)
		} else {
			break
		}

		time.Sleep(3 * time.Hour)
	}

	var videosBuffer []string
	for {
		mostRecentUpload, err := p.client.FetchLatestVideoByChannel(p.ChannelID)

		if err != nil {
		}

		if !contains(videosBuffer, mostRecentUpload) {
			c <- mostRecentUpload
			videosBuffer = append(videosBuffer, mostRecentUpload)
		}

		if len(videosBuffer) == 50 {
			videosBuffer = make([]string, 0)
		}

		time.Sleep(3 * time.Hour)
	}
}

func (p *YoutubePublisher) GetPublisherID() string {
	return p.ChannelID
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
