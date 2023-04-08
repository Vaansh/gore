package youtube

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	ApiKey = "AIzaSyDXCuguEKvISldv2uVWXG0itvKRFzlbueU"
)

type Client struct {
	c      *http.Client
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		c:      &http.Client{},
		apiKey: apiKey,
	}
}

func (c *Client) FetchVideo(videoURL string) (map[string]interface{}, error) {
	resp, err := http.Get(videoURL)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) FetchLatestVideoByChannel(channelID string) string {
	paginatedVideos := PaginatedVideosAPI(channelID)
	resp, err := http.Get(paginatedVideos.String())
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	item := data["items"].([]interface{})[0]
	itemMap := item.(map[string]interface{})
	videoID := itemMap["id"].(map[string]interface{})["videoId"].(string)

	return WatchVideoURL(videoID).String()
}

func WatchVideoURL(videoID string) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "www.youtube.com",
		Path:       "/watch/",
		ForceQuery: true,
		RawQuery:   fmt.Sprintf("v=%s", videoID),
	}
}

func PaginatedVideosAPI(channelId string) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "www.googleapis.com",
		Path:       "/youtube/v3/search",
		ForceQuery: true,
		RawQuery:   fmt.Sprintf("key=%s&channelId=%s&part=snippet,id&order=date&maxResults=25", ApiKey, channelId),
	}
}
