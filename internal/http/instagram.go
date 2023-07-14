package http

import (
	"errors"
	"fmt"
	fb "github.com/huandu/facebook/v2"
	"log"
	"time"
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

func (c *InstagramClient) UploadReel(videoUrl, caption string) bool {
	fmt.Println(videoUrl)
	containerId, err := c.createReelsContainer(videoUrl, caption)
	if err != nil {
		return false
	}

	fmt.Println(containerId)

	ok := c.backoffUntilContainerReady(containerId)
	if !ok {
		return false
	}

	ok = c.uploadReelsByContainer(containerId)
	if !ok {
		return false
	}

	return ok
}

func (c *InstagramClient) createReelsContainer(videoUrl, caption string) (string, error) {
	res, err := fb.Post(c.userId+"/media", fb.Params{
		"media_type":   "REELS",
		"video_url":    videoUrl,
		"caption":      caption,
		"access_token": c.accessToken,
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	containerId, ok := res["id"]
	if !ok {
		log.Println(res)
		return "", errors.New("invalid response")
	}

	return containerId.(string), nil
}

func (c *InstagramClient) uploadReelsByContainer(containerId string) bool {
	res, err := fb.Post(c.userId+"/media_publish", fb.Params{
		"creation_id":  containerId,
		"access_token": c.accessToken,
	})

	if err != nil {
		log.Println(err)
		return false
	}

	_, ok := res["id"]
	if !ok {
		log.Println(err)
		return false
	}

	return true
}

func (c *InstagramClient) backoffUntilContainerReady(containerId string) bool {
	for i := 0; i <= 20; i++ {
		res, err := fb.Get(containerId, fb.Params{
			"fields":       "status_code",
			"access_token": c.accessToken,
		})

		if err != nil {
			log.Println(err)
			return false
		}

		statusCode, ok := res["status_code"]
		if !ok {
			log.Println(res)
			return false
		}

		if statusCode == "FINISHED" {
			return true
		}

		time.Sleep(1 * time.Minute)
	}
	return false
}
