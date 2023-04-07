package main

import (
	"pubsub/internal"
	"pubsub/internal/platform"
)

func main() {
	ChannelID := "UCfeMEuhdUtxtaUMNSvxq_Xg"

	tm := internal.NewTaskManager()

	channels := []string{ChannelID}
	platforms := []platform.PlatformName{platform.PLATFORM}

	err := tm.AddTask(channels, platforms, "", platform.PLATFORM)
	if err != nil {
		return
	}

	tm.RunAll()
}
