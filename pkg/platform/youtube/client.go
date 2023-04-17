package youtube

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	ApiKey = "AIzaSyDXCuguEKvISldv2uVWXG0itvKRFzlbueU"
)

type ChannelShortsListResponse struct {
	Kind  string `json:"kind"`
	Etag  string `json:"etag"`
	Items []struct {
		Kind   string `json:"kind"`
		Etag   string `json:"etag"`
		ID     string `json:"id"`
		Shorts []struct {
			VideoID    string `json:"videoId"`
			Title      string `json:"title"`
			Thumbnails []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"thumbnails"`
			ViewCount int `json:"viewCount"`
		} `json:"shorts"`
		NextPageToken string `json:"nextPageToken"`
	} `json:"items"`
}

type Client struct {
	client *http.Client
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		client: &http.Client{},
		apiKey: apiKey,
	}
}

func (c *Client) FetchVideo(videoURL string) (map[string]interface{}, error) {
	resp, err := http.Get(videoURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Client) FetchLatestVideoByChannel(channelID string) (string, error) {
	paginatedVideos := PaginatedVideosAPI(channelID)
	resp, err := http.Get(paginatedVideos.String())
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	item := data["items"].([]interface{})[0]
	itemMap := item.(map[string]interface{})
	videoID := itemMap["id"].(map[string]interface{})["videoId"].(string)

	return videoID, nil
}

func (c *Client) FetchShortsByChannel(channelId string) ([]string, string, error) {
	paginatedShorts := PaginatedShortsAPI(channelId)
	fmt.Println(paginatedShorts)
	resp, err := http.Get(paginatedShorts.String())
	if err != nil {
		fmt.Println(err)
		return nil, "", err
	}

	defer resp.Body.Close()

	var data ChannelShortsListResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
		return nil, "", err
	}

	var videoIds []string
	var nextPageToken string

	for _, item := range data.Items {
		for _, short := range item.Shorts {
			videoIds = append(videoIds, short.VideoID)
		}
		nextPageToken = item.NextPageToken
	}

	return videoIds, nextPageToken, nil
}

func (c *Client) FetchChannelName(channelId string) (string, error) {
	paginatedVideos := ChannelInfoAPI(channelId)
	resp, err := http.Get(paginatedVideos.String())
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	item := data["items"].([]interface{})[0]
	itemMap := item.(map[string]interface{})
	channelName := itemMap["snippet"].(map[string]interface{})["title"].(string)
	return channelName, nil
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

func ChannelInfoAPI(channelId string) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "www.googleapis.com",
		Path:       "/youtube/v3/channels",
		ForceQuery: true,
		RawQuery:   fmt.Sprintf("key=%s&id=%s&part=snippet,id", ApiKey, channelId),
	}
}

func PaginatedShortsAPI(videoId string) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "yt.lemnoslife.com",
		Path:       "channels",
		ForceQuery: true,
		RawQuery:   fmt.Sprintf("part=shorts&id=%s", videoId),
	}
}
