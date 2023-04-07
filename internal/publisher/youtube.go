package publisher

import (
	"pubsub/pkg/client"
	"time"
)

type YoutubeProducer struct {
	ChannelID string
	client    *client.YoutubeClient
}

func NewYoutubeProducer(channelID string) *YoutubeProducer {
	return &YoutubeProducer{
		ChannelID: channelID,
		client:    client.NewYoutubeClient(client.ApiKey),
	}
}

func (p *YoutubeProducer) PublishTo(c chan<- string) {
	videos := p.client.FetchVideosByChannel(p.ChannelID)
	for _, link := range videos {
		c <- link
	}

	for {
		latestVideoLink := p.client.FetchLatestVideoByChannel(p.ChannelID)

		if !contains(videos, latestVideoLink) {
			c <- latestVideoLink
			videos = append(videos, latestVideoLink)
		}

		time.Sleep(3 * time.Hour)
	}
}

func (p *YoutubeProducer) GetPublisherID() string {
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
