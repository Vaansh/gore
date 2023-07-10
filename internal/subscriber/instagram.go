package subscriber

import (
	"fmt"
	"github.com/Vaansh/gore/internal/platform"
	"github.com/Vaansh/gore/internal/util"
	"time"
)

type InstagramSubscriber struct {
	InstagramID string
}

func NewInstagramSubscriber(InstagramID string) *InstagramSubscriber {
	return &InstagramSubscriber{InstagramID: InstagramID}
}

func (s InstagramSubscriber) SubscribeTo(c <-chan string, name platform.Name) {
	if name == platform.YOUTUBE {
		for videoId := range c {
			fmt.Println(videoId)
			util.SaveYoutubeVideo(videoId)
			// TODO: Client posting logic
			time.Sleep(10 * time.Second)
			util.Delete(videoId)
		}
	}
}

func (s InstagramSubscriber) GetSubscriberID() string {
	return s.InstagramID
}
