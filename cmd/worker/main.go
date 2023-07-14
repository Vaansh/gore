package main

import (
	"fmt"
	"github.com/Vaansh/gore/internal"
	"github.com/Vaansh/gore/internal/database"
	"github.com/Vaansh/gore/internal/model"
	"github.com/Vaansh/gore/internal/platform"
	"os"
)

func main() {
	db, err := database.ConnectDb()
	if err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	ChannelID := "UCQ4zIVlfhsmvds7WuKeL2Bw"
	tm := internal.NewTaskPool()

	channels := []string{ChannelID}
	platforms := []platform.Name{platform.YOUTUBE}

	subscriberId := "pewdiepie_exe"
	subPlatform := platform.INSTAGRAM
	hashtags := "#pewdiepie #memes #dankmemes #meme #funny #memesdaily #dank #funnymemes #pewdiepiememes #memereview #pewds #spicymemes #dankmeme #lwiay #reels #reelsinstagram #instagram #explore #viral #trending #tiktok #shorts #youtube #fyp #gamer #dailymemes #gaming #mrbeast"
	err = tm.AddTask(channels, platforms, subscriberId, subPlatform, model.MetaData{IgPostTags: hashtags}, *database.NewUserRepository(db, subscriberId, subPlatform))
	if err != nil {
		return
	}

	tm.RunAll()
}
