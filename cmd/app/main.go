package main

import (
	"fmt"
	"github.com/Vaansh/gore/pkg/platform/youtube"
)

const (
	ApiKey = "AIzaSyDXCuguEKvISldv2uVWXG0itvKRFzlbueU"
)

func main() {
	c := youtube.Client{}

	//if name, err := c.FetchChannelName("UCfeMEuhdUtxtaUMNSvxq_Xg"); err == nil {
	//	fmt.Println(name)
	//}

	if name, err := c.FetchShortsByChannel("UC5O114-PQNYkurlTg6hekZw"); err == nil {
		fmt.Println(name)
	}

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

	//videoID := "dCcuIFO_SF8"
	//client := ytdownloader.Client{}
	//
	//videoLink, err := client.GetVideo(videoID)
	//if err != nil {
	//	panic(err)
	//}
	//
	//formats := videoLink.Formats.WithAudioChannels()
	//stream, _, err := client.GetStream(videoLink, &formats[0])
	//if err != nil {
	//	panic(err)
	//}
	//
	//file, err := os.Create("video.mp4")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//
	//_, err = io.Copy(file, stream)
	//if err != nil {
	//	panic(err)
	//}
	//
	//insta := goinsta.New("mrbeast.zip", "")
	//
	//if err := insta.Login(); err != nil {
	//	panic(err)
	//}
	//
	//file, err = os.Open("video.mp4")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer file.Close()
	//
	//reader := io.Reader(file)
	//
	//video, err := io.ReadAll(reader)
	//if err != nil {
	//	return // err
	//}
	//
	//var following = insta.Account.Following("", goinsta.DefaultOrder)
	//
	//_, err = insta.Upload(&goinsta.UploadOptions{
	//	File:    bytes.NewReader(video),
	//	Caption: "capt",
	//})
	//
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(len(following.Users))
}

//Mindc-blowing stunts and insane challenges 🤯
//
//Follow @mrbeast.reels for non-stop action 🔥
//
//Credit: @originalposter 💯
//
//#mrbeast #beastgang #reels #challengeaccepted
