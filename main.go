package main

import (
	"pubsub/internal"
	"pubsub/internal/lifecycle"
)

func main() {
	ChannelID := "UCfeMEuhdUtxtaUMNSvxq_Xg"

	tm := lifecycle.NewTaskManager()

	channels := []string{ChannelID}
	platforms := []internal.PlatformName{internal.PLATFORM}

	err := tm.AddTask(channels, platforms, "", internal.PLATFORM)
	if err != nil {
		return
	}

	tm.RunAll()
}
