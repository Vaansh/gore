package main

func main() {
	ChannelID := "UCfeMEuhdUtxtaUMNSvxq_Xg"

	tm := NewTaskManager()

	channels := []string{ChannelID}
	platforms := []PlatformName{PLATFORM}

	err := tm.AddTask(channels, platforms, "", PLATFORM)
	if err != nil {
		return
	}

	tm.RunAll()
}
