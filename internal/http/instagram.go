package http

import (
	"fmt"
	"time"

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

func (c *InstagramClient) UploadReel(videoUrl, coverUrl, caption string) error {
	containerId, err := c.createReelsContainer(videoUrl, coverUrl, caption)
	if err != nil {
		return err
	}

	err = c.backoffUntilContainerReady(containerId)
	if err != nil {
		return err
	}

	err = c.uploadReelsByContainer(containerId)
	if err != nil {
		return err
	}

	return nil
}

func (c *InstagramClient) createReelsContainer(videoUrl, coverUrl, caption string) (string, error) {
	var res fb.Result
	var err error

	if coverUrl == "" {
		res, err = fb.Post(c.userId+"/media", fb.Params{
			"media_type":   "REELS",
			"video_url":    videoUrl,
			"caption":      caption,
			"access_token": c.accessToken,
		})
	} else {
		res, err = fb.Post(c.userId+"/media", fb.Params{
			"media_type":   "REELS",
			"video_url":    videoUrl,
			"caption":      caption,
			"cover_url":    coverUrl,
			"access_token": c.accessToken,
		})
	}

	if err != nil {
		return "", err
	}

	containerId, ok := res["id"]
	if !ok {
		return "", fmt.Errorf("no 'id' field found for videoUrl (%s) response received: %s", videoUrl, res)
	}

	return containerId.(string), nil
}

func (c *InstagramClient) uploadReelsByContainer(containerId string) error {
	res, err := fb.Post(c.userId+"/media_publish", fb.Params{
		"creation_id":  containerId,
		"access_token": c.accessToken,
	})

	if err != nil {
		return err
	}

	_, ok := res["id"]
	if !ok {
		return fmt.Errorf("no 'id' field found for containerId (%s) response received: %s", containerId, res)
	}

	return nil
}

func (c *InstagramClient) backoffUntilContainerReady(containerId string) error {
	for i := 0; i <= 20; i++ {
		res, err := fb.Get(containerId, fb.Params{
			"fields":       "status_code",
			"access_token": c.accessToken,
		})

		if err != nil {
			return err
		}

		statusCode, ok := res["status_code"]
		if !ok {
			return fmt.Errorf("no 'status_code' field found for %s response received: %s", containerId, res)
		}

		if statusCode == "FINISHED" {
			return nil
		}

		time.Sleep(1 * time.Minute)
	}
	return fmt.Errorf("timeout error")
}
