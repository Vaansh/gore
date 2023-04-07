package main

func main() {
	//https://siongui.github.io/2018/02/14/go-get-photo-video-in-instagram-post/
	ChannelID := "UCfeMEuhdUtxtaUMNSvxq_Xg"
	//channel := make(chan string)
	//producer := NewYoutubeProducer(ChannelID)
	//consumer := NewInstagramConsumer("")
	//
	//go producer.ProduceOn(channel)
	//consumer.ConsumeOn(channel)
	//

	//insta := goinsta.New("interestingasfuckk", "ASD123___")
	//
	//if err := insta.Login(); err != nil {
	//	panic(err)
	//}

	tm := NewTaskManager()
	pid := []string{ChannelID}
	plats := []PlatformName{PLATFORM}
	err := tm.AddTask(pid, plats, "", PLATFORM)
	if err != nil {
		return
	}

	tm.RunAll()
}
