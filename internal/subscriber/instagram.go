package subscriber

import (
	"github.com/Vaansh/gore"
	"github.com/Vaansh/gore/internal/database"
	"github.com/Vaansh/gore/internal/gcloud"
	"github.com/Vaansh/gore/internal/http"
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/util"
	"log"
	"strings"
	"time"
)

type InstagramSubscriber struct {
	instagramId string
	hashtags    string
	repository  database.UserRepository
	client      *http.InstagramClient
}

func NewInstagramSubscriber(instagramId string, metadata model.MetaData, repository database.UserRepository) *InstagramSubscriber {
	return &InstagramSubscriber{
		instagramId: instagramId,
		hashtags:    metadata.IgPostTags,
		repository:  repository,
		client:      http.NewInstagramClient(metadata.IgUserId, metadata.IgLongLivedAccessToken),
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
			if sourcePlatform == go_pubsub.YOUTUBE {
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
			ok = s.client.UploadReel(fileUrl, util.GenerateInstagramCaption(caption, author, s.hashtags, strings.ToUpper(sourcePlatform.String())))
			if !ok {
				break
			}

			err = fileHandler.DeleteFromBucket(fileName)
			if err != nil {
				log.Println(err)
			}

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
	return go_pubsub.INSTAGRAM.String() + "_" + s.instagramId
}
