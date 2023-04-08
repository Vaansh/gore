package main

import (
	"github.com/kkdai/youtube/v2"
	"io"
	"os"
)

func main() {
	//ChannelID := "UCfeMEuhdUtxtaUMNSvxq_Xg"

	//tm := app.NewTaskManager()
	//
	//channels := []string{ChannelID}
	//platforms := []pkg.PlatformName{pkg.PLATFORM}
	//
	//err := tm.AddTask(channels, platforms, "", pkg.PLATFORM)
	//if err != nil {
	//	return
	//}
	//
	//tm.RunAll()

	//insta := goinsta.New("USERNAME", "PASSWORD")
	//
	//// Only call Login the first time you login. Next time import your config
	//if err := insta.Login(); err != nil {
	//	panic(err)
	//}
	//
	//item, err := insta.Upload(
	//	&goinsta.UploadOptions{
	//		File:    videoFile,
	//		Caption: "awesome! :)",
	//	},
	//)

	videoID := "dCcuIFO_SF8"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels()
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}
