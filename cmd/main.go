package main

import (
	"github.com/Vaansh/gore/internal/lifecycle"
	"github.com/Vaansh/gore/internal/platform"
)

func main() {
	ChannelID := "UCfeMEuhdUtxtaUMNSvxq_Xg"

	tm := lifecycle.NewTaskManager()

	channels := []string{ChannelID}
	platforms := []platform.PlatformName{platform.PLATFORM}

	err := tm.AddTask(channels, platforms, "", platform.PLATFORM)
	if err != nil {
		return
	}

	tm.RunAll()
}
