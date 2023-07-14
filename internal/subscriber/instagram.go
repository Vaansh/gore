package subscriber

import (
	"fmt"
	"github.com/Vaansh/gore/internal/database"
	"github.com/Vaansh/gore/internal/gcloud"
	"github.com/Vaansh/gore/internal/http"
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/platform"
	"github.com/Vaansh/gore/internal/util"
	"strings"
	"time"
)

type InstagramSubscriber struct {
	instagramId string
	metadata    model.MetaData
	repository  database.UserRepository
	client      *http.InstagramClient
}

func NewInstagramSubscriber(instagramId string, metadata model.MetaData, repository database.UserRepository) *InstagramSubscriber {
	userId := util.Getenv("IG_USER_ID", true)
	accessToken := util.Getenv("LONG_LIVED_ACCESS_TOKEN", true)
	return &InstagramSubscriber{
		instagramId: instagramId,
		metadata:    metadata,
		repository:  repository,
		client:      http.NewInstagramClient(userId, accessToken),
	}
}

func (s *InstagramSubscriber) SubscribeTo(c <-chan model.Post) {
	fileHandler := gcloud.NewStorageHandler()
	for post := range c {
		postId, author, sourcePlatform, caption := post.GetParams()
		exists, err := s.repository.CheckIfRecordExists(s.getTableName(), &post)
		if err != nil {
		}

		if !exists {
			ok := false
			if sourcePlatform == platform.YOUTUBE {
				ok = fileHandler.SaveYoutubeVideo(postId)
			}

			if !ok {
				break
			}

			fileName := sourcePlatform.String() + "_" + postId + ".mp4"
			ok = fileHandler.UploadToBucket(fileName)
			if !ok {
				break
			}

			fileUrl := fileHandler.GetFileUrl(fileName)
			ok = s.client.UploadReel(fileUrl, util.GenerateInstagramCaption(caption, author, s.metadata.IgPostTags, strings.ToUpper(sourcePlatform.String())))
			if !ok {
				fmt.Println("3")
				break
			}

			fmt.Println("Persisting record to db")
			err = s.repository.AddRecord(s.getTableName(), &post)
			if err != nil {
				break
			}

			gcloud.Delete(fileName)
			time.Sleep(30 * time.Minute)
		}
	}
}

func (s *InstagramSubscriber) GetSubscriberId() string {
	return s.instagramId
}

func (s *InstagramSubscriber) getTableName() string {
	return platform.INSTAGRAM.String() + "_" + s.instagramId
}
