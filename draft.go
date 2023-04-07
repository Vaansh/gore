package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"time"
//)
//
//const (
//	apiKey        = "AIzaSyDXCuguEKvISldv2uVWXG0itvKRFzlbueU"
//	channelID     = "UCfeMEuhdUtxtaUMNSvxq_Xg"
//	baseVideoURL  = "https://www.youtube.com/watch?v="
//	baseSearchURL = "https://www.googleapis.com/youtube/v3/search?"
//	maxResults    = 25
//	waitTime      = 6 * time.Hour // Time interval between each check for new videos
//)
//
//func main() {
//	videoIDs := make(map[string]bool)
//	lastCheckTime := time.Now().Add(-waitTime) // Set initial time as if a check has been made
//	for {
//		// Check if it's time to check for new videos
//		if time.Now().Sub(lastCheckTime) >= waitTime {
//			newVideoIDs := getNewVideoIDs()
//			for videoID := range newVideoIDs {
//				if !videoIDs[videoID] {
//					videoURL := baseVideoURL + videoID
//					fmt.Println(videoURL)
//					videoIDs[videoID] = true
//				}
//			}
//			lastCheckTime = time.Now()
//		}
//
//		time.Sleep(time.Minute)
//	}
//}
//
//func getNewVideoIDs() map[string]bool {
//	newVideoIDs := make(map[string]bool)
//	pageToken := ""
//	for {
//		url := fmt.Sprintf("%skey=%s&channelId=%s&part=snippet,id&order=date&maxResults=%d&pageToken=%s",
//			baseSearchURL, apiKey, channelID, maxResults, pageToken)
//		resp, err := http.Get(url)
//		if err != nil {
//			fmt.Println("Error getting videos:", err)
//			return newVideoIDs
//		}
//		defer resp.Body.Close()
//
//		var searchResult struct {
//			Items []struct {
//				ID struct {
//					Kind    string `json:"kind"`
//					VideoID string `json:"videoId"`
//				} `json:"id"`
//			} `json:"items"`
//			NextPageToken string `json:"nextPageToken"`
//		}
//
//		err = json.NewDecoder(resp.Body).Decode(&searchResult)
//		if err != nil {
//			fmt.Println("Error decoding search result:", err)
//			return newVideoIDs
//		}
//
//		for _, item := range searchResult.Items {
//			if item.ID.Kind == "youtube#video" {
//				newVideoIDs[item.ID.VideoID] = true
//			}
//		}
//
//		if searchResult.NextPageToken == "" {
//			break
//		}
//		pageToken = searchResult.NextPageToken
//	}
//
//	return newVideoIDs
//}
