package main

import (
	"fmt"
	"pubsub/youtube"
	"time"
)

//iterate and print the link over every single youtube video from a channel
//then keep monitoring every few hours and print when the next video is uploaded

const (
	ApiKey = "AIzaSyDXCuguEKvISldv2uVWXG0itvKRFzlbueU"
)

func main() {
	// r: DestinationPost [eg: InstagramPost]
	// c chan SourcePost
	// receiver: processor
	// --------------------
	//c := make(chan InstagramPost)

	cl := youtube.NewClient(ApiKey)
	cl.NextVideo()

	//go sender(c)
	//go sender(c)
	//go receiver(c)
	//select {}

	//channelID := "UCfeMEuhdUtxtaUMNSvxq_Xg"

	//// Set up channel for new video links
	//newVideos := make(chan string)
	//
	//// Start a goroutine to receive and handle new video links
	//go func() {
	//	for link := range newVideos {
	//		fmt.Printf("New video uploaded: %s\n", link)
	//		time.Sleep(4 * time.Second)
	//	}
	//}()

	// Get all video links from the channel
	//videoLinks := FetchVideosByChannel(channelID)
	//
	//// Print all video links
	//for _, link := range videoLinks {
	//	newVideos <- link
	//}
	//
	//// Start monitoring for new videos every few hours
	//for {
	//	// Get the latest video link from the channel
	//	latestVideoLink := FetchLatestVideoByChannel(channelID)
	//
	//	// Publish the latest video link to the channel if it's not in the list of existing video links
	//	if !contains(videoLinks, latestVideoLink) {
	//		newVideos <- latestVideoLink
	//		videoLinks = append(videoLinks, latestVideoLink)
	//	}
	//
	//	// Wait for a few hours before checking again
	//	time.Sleep(3 * time.Hour)
	//}
}

//
//func FetchVideosByChannel(channelID string) []string {
//	baseVideoURL := "https://www.youtube.com/watch?v="
//	baseSearchURL := "https://www.googleapis.com/youtube/v3/search?"
//
//	firstURL := fmt.Sprintf("%skey=%s&channelId=%s&part=snippet,id&order=date&maxResults=25", baseSearchURL, ApiKey, channelID)
//
//	var videoLinks []string
//	url := firstURL
//	for {
//		resp, err := http.Get(url)
//		if err != nil {
//			fmt.Println(err)
//			break
//		}
//
//		defer resp.Body.Close()
//
//		var data map[string]interface{}
//		err = json.NewDecoder(resp.Body).Decode(&data)
//		if err != nil {
//			fmt.Println(err)
//			break
//		}
//
//		for _, item := range data["items"].([]interface{}) {
//			itemMap := item.(map[string]interface{})
//			if itemMap["id"].(map[string]interface{})["kind"] == "youtube#video" {
//				videoID := itemMap["id"].(map[string]interface{})["videoId"].(string)
//				videoLinks = append(videoLinks, baseVideoURL+videoID)
//			}
//		}
//
//		if nextPageToken, ok := data["nextPageToken"].(string); ok {
//			url = fmt.Sprintf("%s&pageToken=%s", firstURL, nextPageToken)
//		} else {
//			break
//		}
//	}
//
//	return videoLinks
//}
//
//func FetchLatestVideoByChannel(channelID string) string {
//	baseVideoURL := "https://www.youtube.com/watch?v="
//	baseSearchURL := "https://www.googleapis.com/youtube/v3/search?"
//
//	latestURL := fmt.Sprintf("%skey=%s&channelId=%s&part=snippet,id&order=date&maxResults=1", baseSearchURL, ApiKey, channelID)
//
//	resp, err := http.Get(latestURL)
//	if err != nil {
//		fmt.Println(err)
//		return ""
//	}
//
//	defer resp.Body.Close()
//
//	var data map[string]interface{}
//	err = json.NewDecoder(resp.Body).Decode(&data)
//	if err != nil {
//		fmt.Println(err)
//		return ""
//	}
//
//	item := data["items"].([]interface{})[0]
//	itemMap := item.(map[string]interface{})
//	videoID := itemMap["id"].(map[string]interface{})["videoId"].(string)
//
//	return baseVideoURL + videoID
//}
//
//func contains(slice []string, item string) bool {
//	for _, s := range slice {
//		if s == item {
//			return true
//		}
//	}
//	return false
//}

func sender(c chan InstagramPost) {
	//provider := youtube.YoutubeProvider{}
	for {
		//provider.NextVideo()
		p := NewInstagramPost("", "", "")
		c <- p
	}
}

func receiver(c chan InstagramPost) {
	for r := range c {
		fmt.Println("Received:", r)
		time.Sleep(5 * time.Second)
	}

	//for {
	//	select {
	//	case r := <-c:
	//		{
	//			fmt.Println("Received:", r)
	//			time.Sleep(5 * time.Second)
	//		}
	//	default:
	//	}
	//}
}
