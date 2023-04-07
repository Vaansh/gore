package youtube

import (
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
	//TODO: fetch 25 from client, move loop logic here instead
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
