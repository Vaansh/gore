package youtube

import (
	"fmt"
	"time"
)

type Publisher struct {
	ChannelID string
	client    *Client
}

func NewPublisher(channelID string) *Publisher {
	return &Publisher{
		ChannelID: channelID,
		client:    NewClient(ApiKey),
	}
}

func (p *Publisher) PublishTo(c chan<- string) {
	firstURL := PaginatedVideosAPI(p.ChannelID)
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
				videoLink := WatchVideoURL(videoID).String()
				c <- videoLink
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
		mostRecentUpload := p.client.FetchLatestVideoByChannel(p.ChannelID)
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

func (p *Publisher) GetPublisherID() string {
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
