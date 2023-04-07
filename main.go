package main

import (
	"pubsub/internal"
)

func main() {
	ChannelID := "UCfeMEuhdUtxtaUMNSvxq_Xg"

	tm := internal.NewTaskManager()
	pid := []string{ChannelID}
	plats := []internal.PlatformName{internal.PLATFORM}

	err := tm.AddTask(pid, plats, "", internal.PLATFORM)
	if err != nil {
		return
	}

	tm.RunAll()
}
