package main

import (
	"fmt"
	"github.com/Vaansh/gore/internal"
	"github.com/Vaansh/gore/internal/config"
	"github.com/Vaansh/gore/internal/platform"
	"os"
)

func main() {
	db, err := config.ConnectToDB()
	if err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()
	//c := http.NewYoutubeClient(ApiKey)

	//if name, err := c.FetchChannelName("UCfeMEuhdUtxtaUMNSvxq_Xg"); err == nil {
	//	fmt.Println(name)
	//}

	//FANTANO: UCfpcfju9rBs5o_xQLXmLQHQ
	//SOMEORDINARYGAMERS: UCtMVHI3AJD4Qk4hcbZnI9ZQ
	//name, err := c.FetchLatestShortByChannel("UCfpcfju9rBs5o_xQLXmLQHQ")
	//if err != nil {
	//	fmt.Println("err..")
	//}
	//
	//fmt.Println(name)

	ChannelID := "UCfpcfju9rBs5o_xQLXmLQHQ"
	// Logic1: fetch and exec ig to yt:
	//ChannelID := "UCtMVHI3AJD4Qk4hcbZnI9ZQ"
	tm := internal.NewTaskPool()

	channels := []string{ChannelID}
	platforms := []platform.Name{platform.YOUTUBE}

	err = tm.AddTask(channels, platforms, "", platform.INSTAGRAM)
	if err != nil {
		return
	}

	tm.RunAll()
	// -- Logic1 --

	// Logic2: downlaod yt vid
	//videoId := "dCcuIFO_SF8"
	// -- Logic2 --

	// Logic3: delete file

	// -- Logic3 --
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

	// Logic4: read env files
	//err := godotenv.Load(".env")
	//
	//if err != nil {
	//	log.Fatalf("Error loading environment variables file")
	//}
	//
	//name := os.Getenv("ig_user_id")
	//other := os.Getenv("ACCESS_TOKEN")
	//
	//fmt.Println(name)
	//fmt.Println(other)
	// -- Logic4 --

}

//Mindc-blowing stunts and insane challenges 🤯
//
//Follow @mrbeast.reels for non-stop action 🔥
//
//Credit: @originalposter 💯
//
//#mrbeast #beastgang #reels #challengeaccepted
