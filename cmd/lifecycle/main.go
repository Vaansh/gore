package main

import (
	"github.com/Vaansh/gore/internal/lifecycle"
	"github.com/Vaansh/gore/pkg"
)

func main() {
	ChannelID := "UCfeMEuhdUtxtaUMNSvxq_Xg"

	tm := lifecycle.NewTaskManager()

	channels := []string{ChannelID}
	platforms := []pkg.PlatformName{pkg.PLATFORM}

	err := tm.AddTask(channels, platforms, "", pkg.PLATFORM)
	if err != nil {
		return
	}

	tm.RunAll()
}
