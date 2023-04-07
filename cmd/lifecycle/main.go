package main

import (
	"github.com/Vaansh/gore/internal"
	"github.com/Vaansh/gore/pkg/platform"
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
