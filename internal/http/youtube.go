package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/platform"
	"net/http"
	"net/url"
)

type YoutubeClient struct {
	apiKey string
}

func NewYoutubeClient(apiKey string) *YoutubeClient {
	return &YoutubeClient{apiKey: apiKey}
}

// Shorts
func (c *YoutubeClient) FetchLatestShortByChannel(channelId string) (model.Post, error) {
	paginatedShorts := PaginatedShortsAPI(channelId)

	resp, err := http.Get(paginatedShorts.String())
	if err != nil {
		fmt.Println(err)
		return model.Post{}, err
	}

	defer resp.Body.Close()

	var response shortsListByChannelResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return model.Post{}, err
	}

	if len(response.Items) < 1 || len(response.Items[0].Shorts) < 1 {
		return model.Post{}, errors.New("invalid response")
	}

	author, err := c.FetchChannelName(channelId)
	if err != nil {
	}
	return *model.NewPost(response.Items[0].Shorts[0].VideoID, response.Items[0].Shorts[0].Title, author, channelId, platform.YOUTUBE), nil
}

func (c *YoutubeClient) FetchPaginatedShortsByChannel(channelId string) ([]model.Post, string, error) {
	paginatedShorts := PaginatedShortsAPI(channelId)
	resp, err := http.Get(paginatedShorts.String())
	if err != nil {
		fmt.Println(err)
		return nil, "", err
	}

	defer resp.Body.Close()

	var response shortsListByChannelResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return nil, "", err
	}

	var posts []model.Post
	var nextPageToken string

	author, err := c.FetchChannelName(channelId)
	if err != nil {
	}

	for _, item := range response.Items {
		for _, short := range item.Shorts {
			posts = append(posts, *model.NewPost(short.VideoID, short.Title, author, channelId, platform.YOUTUBE))
		}
		nextPageToken = item.NextPageToken
	}

	return posts, nextPageToken, nil
}

// Videos
func (c *YoutubeClient) FetchVideo(videoURL string) (map[string]interface{}, error) {
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

func (c *YoutubeClient) FetchLatestVideoByChannel(channelID string) (string, error) {
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

// Channels
func (c *YoutubeClient) FetchChannelName(channelId string) (string, error) {
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

func (c *YoutubeClient) PaginatedVideosAPI(channelId string) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "www.googleapis.com",
		Path:       "/youtube/v3/search",
		ForceQuery: true,
		RawQuery:   fmt.Sprintf("key=%s&channelId=%s&part=snippet,id&order=date&maxResults=25&type=video", c.apiKey, channelId),
	}
}

func (c *YoutubeClient) ChannelInfoAPI(channelId string) *url.URL {
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
		RawQuery:   fmt.Sprintf("part=shorts&id=%s", videoId),
	}
}

type thumbnails struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type short struct {
	VideoID    string       `json:"videoId"`
	Title      string       `json:"title"`
	Thumbnails []thumbnails `json:"thumbnails"`
	ViewCount  int          `json:"viewCount"`
}

type item struct {
	Kind          string  `json:"kind"`
	Etag          string  `json:"etag"`
	ID            string  `json:"id"`
	Shorts        []short `json:"shorts"`
	NextPageToken string  `json:"nextPageToken"`
}

type shortsListByChannelResponse struct {
	Kind  string `json:"kind"`
	Etag  string `json:"etag"`
	Items []item `json:"items"`
}
