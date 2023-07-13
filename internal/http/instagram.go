package http

import (
	"errors"
	fb "github.com/huandu/facebook/v2"
)

type InstagramClient struct {
	userId      string
	accessToken string
}

func NewInstagramClient(userId, accessToken string) *InstagramClient {
	return &InstagramClient{
		userId:      userId,
		accessToken: accessToken,
	}
}

func (c *InstagramClient) UploadReel(videoUrl string) bool {
	containerId, err := c.createReelsContainer(videoUrl)
	if err != nil {
		//	Log err here
	}

	uploadStatus, err := c.uploadReelsByContainer(containerId)
	if err != nil {
		//	Log err here
	}

	return uploadStatus
}

func (c *InstagramClient) createReelsContainer(videoUrl string) (string, error) {
	res, err := fb.Post(c.userId+"/media", fb.Params{
		"media_type":   "REELS",
		"video_url":    videoUrl,
		"access_token": c.accessToken,
	})

	if err != nil {
		return "", errors.New("invalid request")
	}

	containerId, ok := res["id"]
	if !ok {
		return "", errors.New("invalid response")
	}

	return containerId.(string), nil
}

func (c *InstagramClient) uploadReelsByContainer(containerId string) (bool, error) {
	res, err := fb.Post(c.userId+"/media_publish", fb.Params{
		"creation_id":  containerId,
		"access_token": c.accessToken,
	})

	if err != nil {

	}

	_, ok := res["id"]
	if !ok {
		return false, errors.New("")
	}

	return true, nil
}
