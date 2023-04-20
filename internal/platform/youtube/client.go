package youtube

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: "AIzaSyDXCuguEKvISldv2uVWXG0itvKRFzlbueU",
	}
}

// Videos
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
	paginatedVideos := c.PaginatedVideosAPI(channelID)
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

// Shorts
func (c *Client) FetchLatestShortByChannel(channelId string) (string, error) {
	paginatedShorts := PaginatedShortsAPI(channelId)
	fmt.Println(paginatedShorts)
	resp, err := http.Get(paginatedShorts.String())
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer resp.Body.Close()

	var response ChannelShortsListResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if len(response.Items) < 1 || len(response.Items[0].Shorts) < 1 {
		return "", errors.New("invalid response")
	}

	return response.Items[0].Shorts[0].VideoID, nil
}

func (c *Client) FetchPaginatedShortsByChannel(channelId string) ([]string, string, error) {
	paginatedShorts := PaginatedShortsAPI(channelId)
	fmt.Println(paginatedShorts)
	resp, err := http.Get(paginatedShorts.String())
	if err != nil {
		fmt.Println(err)
		return nil, "", err
	}

	defer resp.Body.Close()

	var response ChannelShortsListResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return nil, "", err
	}

	var videoIds []string
	var nextPageToken string

	for _, item := range response.Items {
		for _, short := range item.Shorts {
			videoIds = append(videoIds, short.VideoID)
		}
		nextPageToken = item.NextPageToken
	}

	return videoIds, nextPageToken, nil
}

// Channels
func (c *Client) FetchChannelName(channelId string) (string, error) {
	paginatedVideos := c.ChannelInfoAPI(channelId)
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

func (c *Client) PaginatedVideosAPI(channelId string) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "www.googleapis.com",
		Path:       "/youtube/v3/search",
		ForceQuery: true,
		RawQuery:   fmt.Sprintf("key=%s&channelId=%s&part=snippet,id&order=date&maxResults=25&type=video", c.apiKey, channelId),
	}
}

func (c *Client) ChannelInfoAPI(channelId string) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "www.googleapis.com",
		Path:       "/youtube/v3/channels",
		ForceQuery: true,
		RawQuery:   fmt.Sprintf("key=%s&id=%s&part=snippet,id", c.apiKey, channelId),
	}
}

func PaginatedShortsAPI(videoId string) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "yt.lemnoslife.com",
		Path:       "channels",
		ForceQuery: true,
		RawQuery:   fmt.Sprintf("part=shorts&id=%s&order=date", videoId),
	}
}
