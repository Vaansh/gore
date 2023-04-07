package main

import (
	client2 "pubsub/pkg/client"
	"time"
)

type Producer interface {
	ProduceOn(c chan<- string)
	ProducerID() string
}

type YoutubeProducer struct {
	ChannelID string
	client    *client2.YoutubeClient
}

func NewYoutubeProducer(channelID string) *YoutubeProducer {
	return &YoutubeProducer{
		ChannelID: channelID,
		client:    client2.NewYoutubeClient(client2.ApiKey),
	}
}

func (p *YoutubeProducer) ProduceOn(c chan<- string) {
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

func (p *YoutubeProducer) ProducerID() string {
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
