package publisher

import (
	"github.com/Vaansh/gore/internal/gcloud"
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

func (p *YoutubePublisher) PublishTo(c chan<- model.Post, quit <-chan struct{}) {
	for {
		posts, nextPageToken, err := p.client.FetchPaginatedShortsByChannel(p.channelId)
		if err != nil {
			gcloud.LogWarning(err.Error())
		}

		for _, post := range posts {
			select {
			case c <- post:
			case <-quit:
				return
			}
		}

		if nextPageToken == "" {
			break
		}

		time.Sleep(10 * time.Second)
	}

	var videosBuffer []string
	for {
		post, err := p.client.FetchLatestShortByChannel(p.channelId)
		if err != nil {
			gcloud.LogWarning(err.Error())
		}

		if !util.Contains(videosBuffer, post.PostId) {
			select {
			case c <- post:
				videosBuffer = append(videosBuffer, post.PostId)
			case <-quit:
				return
			}
		}

		if len(videosBuffer) == 50 {
			videosBuffer = make([]string, 0)
		}

		time.Sleep(10 * time.Second)
	}
}

func (p *YoutubePublisher) GetPublisherId() string {
	return p.channelId
}
