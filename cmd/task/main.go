package main

import (
	"fmt"
	"github.com/Vaansh/gore/internal/http"
)

const (
	ApiKey = "AIzaSyDXCuguEKvISldv2uVWXG0itvKRFzlbueU"
)

func main() {
	c := http.NewClient(ApiKey)

	//if name, err := c.FetchChannelName("UCfeMEuhdUtxtaUMNSvxq_Xg"); err == nil {
	//	fmt.Println(name)
	//}

	//FANTANO: UCfpcfju9rBs5o_xQLXmLQHQ
	//SOMEORDINARYGAMERS: UCtMVHI3AJD4Qk4hcbZnI9ZQ
	name, err := c.FetchLatestShortByChannel("UCfpcfju9rBs5o_xQLXmLQHQ")
	if err != nil {
		fmt.Println("err..")
	}

	fmt.Println(name)

	//ChannelID := "UCfeMEuhdUtxtaUMNSvxq_Xg"

	//tm := task.NewTaskManager()
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
	//http := ytdownloader.YoutubeClient{}
	//
	//videoLink, err := http.GetVideo(videoID)
	//if err != nil {
	//	panic(err)
	//}
	//
	//formats := videoLink.Formats.WithAudioChannels()
	//stream, _, err := http.GetStream(videoLink, &formats[0])
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
