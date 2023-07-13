package subscriber

import (
	"fmt"
	"github.com/Vaansh/gore/internal/database"
	"github.com/Vaansh/gore/internal/http"
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/platform"
	"github.com/Vaansh/gore/internal/util"
	"time"
)

type InstagramSubscriber struct {
	instagramId string
	repository  database.UserRepository
	client      *http.InstagramClient
}

func NewInstagramSubscriber(instagramId string, repository database.UserRepository) *InstagramSubscriber {
	userId := util.Getenv("IG_USER_ID", true)
	accessToken := util.Getenv("LONG_LIVED_ACCESS_TOKEN", true)
	return &InstagramSubscriber{
		instagramId: instagramId,
		repository:  repository,
		client:      http.NewInstagramClient(userId, accessToken),
	}
}

func (s InstagramSubscriber) SubscribeTo(c <-chan model.Post) {
	for post := range c {
		postId, _, sourcePlatform, _, _ := post.GetParams()

		// TODO: move to gcloud storage
		if sourcePlatform == platform.YOUTUBE {
			util.SaveYoutubeVideo(postId)
		}

		exists, err := s.repository.CheckIfRecordExists(s.getTableName(), &post)
		if err != nil {
		}

		if !exists {
			ok := s.client.UploadReel("")
			if !ok {
				break
			}

			fmt.Println("Persisting record to db")
			err = s.repository.AddRecord(s.getTableName(), &post)
			if err != nil {
				break
			}
		}

		util.Delete(postId)
		time.Sleep(10 * time.Second)
	}
}

func (s InstagramSubscriber) GetSubscriberId() string {
	return s.instagramId
}

func (s InstagramSubscriber) getTableName() string {
	return platform.INSTAGRAM.String() + "_" + s.instagramId
}
