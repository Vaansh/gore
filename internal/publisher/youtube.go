package publisher

import (
	"fmt"
	"github.com/Vaansh/gore/internal/http"
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/util"
	"time"
)

type YoutubePublisher struct {
	channelId string
	client    *http.YoutubeClient
}

func NewYoutubePublisher(channelId string) *YoutubePublisher {
	apiKey := util.Getenv("YOUTUBE_API_KEY", true)
	return &YoutubePublisher{
		channelId: channelId,
		client:    http.NewYoutubeClient(apiKey),
	}
}

func (p YoutubePublisher) PublishTo(c chan<- model.Post) {
	fmt.Println("Fetching Paginated shorts")
	for {
		posts, nextPageToken, err := p.client.FetchPaginatedShortsByChannel(p.channelId)
		if err != nil {
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
		post, err := p.client.FetchLatestShortByChannel(p.channelId)

		if err != nil {
		}

		if !util.Contains(videosBuffer, post.PostId) {
			c <- post
			videosBuffer = append(videosBuffer, post.PostId)
		}

		if len(videosBuffer) == 50 {
			videosBuffer = make([]string, 0)
		}

		time.Sleep(10 * time.Second)
	}
}

func (p YoutubePublisher) GetPublisherId() string {
	return p.channelId
}
