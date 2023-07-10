package publisher

import (
	"fmt"
	"github.com/Vaansh/gore/internal/http"
	"github.com/Vaansh/gore/internal/model"
	"time"
)

type YoutubePublisher struct {
	ChannelID string
	client    *http.YoutubeClient
}

func NewYoutubePublisher(channelID string) *YoutubePublisher {
	return &YoutubePublisher{
		ChannelID: channelID,
		client:    http.NewYoutubeClient(),
	}
}

func (p *YoutubePublisher) PublishVideosTo(c chan<- string) {
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

func (p YoutubePublisher) PublishTo(c chan<- model.Post) {
	fmt.Println("Fetching Paginated shorts")
	for {
		posts, nextPageToken, err := p.client.FetchPaginatedShortsByChannel(p.ChannelID)

		if err != nil {
			fmt.Println(err)
			break
		}

		for _, post := range posts {
			c <- post
		}

		if nextPageToken == "" {
			break
		}

		time.Sleep(10 * time.Second)
	}

	fmt.Println("Fetching New shorts")
	var videosBuffer []string
	for {
		post, err := p.client.FetchLatestShortByChannel(p.ChannelID)

		if err != nil {
		}

		if !contains(videosBuffer, post.ID) {
			c <- post
			videosBuffer = append(videosBuffer, post.ID)
		}

		if len(videosBuffer) == 50 {
			videosBuffer = make([]string, 0)
		}

		time.Sleep(10 * time.Second)
	}
}

func (p YoutubePublisher) GetPublisherID() string {
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
