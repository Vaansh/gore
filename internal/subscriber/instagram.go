package subscriber

import (
	"fmt"
	"github.com/Vaansh/gore/internal/database"
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/platform"
	"github.com/Vaansh/gore/internal/util"
	"time"
)

type InstagramSubscriber struct {
	InstagramID string
	repository  database.UserRepository
}

func NewInstagramSubscriber(instagramId string, repository database.UserRepository) *InstagramSubscriber {
	return &InstagramSubscriber{InstagramID: instagramId, repository: repository}
}

func (s InstagramSubscriber) SubscribeTo(c <-chan model.Post) {
	for post := range c {
		fmt.Println(post)
		id, _, sourcePlatform, _ := post.GetParams()
		if sourcePlatform == platform.YOUTUBE {
			util.SaveYoutubeVideo(id)
		}
		// TODO: Client posting logic
		time.Sleep(10 * time.Second)
		util.Delete(id)
	}
}

func (s InstagramSubscriber) GetSubscriberID() string {
	return s.InstagramID
}
