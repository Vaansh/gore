package subscriber

import (
	"github.com/Vaansh/gore/internal"
	"github.com/Vaansh/gore/internal/util"
	"time"
)

type InstagramSubscriber struct {
	InstagramID string
}

func NewSubscriber(InstagramID string) *InstagramSubscriber {
	return &InstagramSubscriber{InstagramID: InstagramID}
}

func (p *InstagramSubscriber) SubscribeTo(c <-chan string, platform internal.PlatformName) {
	if platform == internal.PLATFORM {
		for videoId := range c {
			util.SaveYoutubeVideo(videoId)
			// TODO: Client posting logic
			time.Sleep(10 * time.Second)
			util.Delete(videoId)
		}
	}

}

func (p *InstagramSubscriber) GetSubscriberID() string {
	return p.InstagramID
}
