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
	instagramId string
	repository  database.UserRepository
}

func NewInstagramSubscriber(instagramId string, repository database.UserRepository) *InstagramSubscriber {
	return &InstagramSubscriber{instagramId: instagramId, repository: repository}
}

func (s InstagramSubscriber) SubscribeTo(c <-chan model.Post) {
	for post := range c {
		postId, _, sourcePlatform, _, _ := post.GetParams()

		if sourcePlatform == platform.YOUTUBE {
			util.SaveYoutubeVideo(postId)
		}

		exists, err := s.repository.CheckIfRecordExists(s.getTableName(), &post)
		if err != nil {
		}

		if !exists {
			// TODO: Client posting logic
			fmt.Println("Persisting record to db")
			err = s.repository.AddRecord(s.getTableName(), &post)
			if err != nil {
			}
		}

		util.Delete(postId)
		time.Sleep(10 * time.Second)
	}
}

func (s InstagramSubscriber) GetSubscriberID() string {
	return s.instagramId
}

func (s InstagramSubscriber) getTableName() string {
	return platform.INSTAGRAM.String() + "_" + s.instagramId
}
