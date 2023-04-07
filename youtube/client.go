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

type YoutubeClient struct {
	c      *http.Client
	apiKey string
}

func NewYoutubeClient(apiKey string) *YoutubeClient {
	return &YoutubeClient{
		c:      &http.Client{},
		apiKey: apiKey,
	}
}

func (c *YoutubeClient) FetchVideosByChannel(channelID string) []string {
	firstURL := YoutubePaginatedVideosURL(channelID)
	fmt.Println(firstURL)

	var videoLinks []string
	videoURL := firstURL.String()
	for {
		resp, err := http.Get(videoURL)
		if err != nil {
			fmt.Println(err)
			break
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
			break
		}

		for _, item := range data["items"].([]interface{}) {
			itemMap := item.(map[string]interface{})
			if itemMap["id"].(map[string]interface{})["kind"] == "youtube#video" {
				videoID := itemMap["id"].(map[string]interface{})["videoId"].(string)
				videoLinks = append(videoLinks, YoutubeWatchURL(videoID).String())
			}
		}

		if nextPageToken, ok := data["nextPageToken"].(string); ok {
			videoURL = fmt.Sprintf("%s&pageToken=%s", firstURL, nextPageToken)
		} else {
			break
		}
	}

	return videoLinks
}

func (c *YoutubeClient) FetchLatestVideoByChannel(channelID string) string {
	paginatedVideos := YoutubePaginatedVideosURL(channelID)
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

	return YoutubeWatchURL(videoID).String()
}

func YoutubeWatchURL(videoID string) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "www.youtube.com",
		Path:       "/watch/",
		ForceQuery: true,
		RawQuery:   fmt.Sprintf("v=%s", videoID),
	}
}

func YoutubePaginatedVideosURL(channelId string) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "www.googleapis.com",
		Path:       "/youtube/v3/search",
		ForceQuery: true,
		RawQuery:   fmt.Sprintf("key=%s&channelId=%s&part=snippet,id&order=date&maxResults=25", ApiKey, channelId),
	}
}
