package main

import (
	"fmt"
	"github.com/Vaansh/gore/internal"
	"github.com/Vaansh/gore/internal/database"
	"github.com/Vaansh/gore/internal/platform"
	"os"
)

func main() {
	db, err := database.ConnectDb()
	if err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	//UCfeMEuhdUtxtaUMNSvxq_Xg
	//FANTANO: UCfpcfju9rBs5o_xQLXmLQHQ
	//SOMEORDINARYGAMERS: UCtMVHI3AJD4Qk4hcbZnI9ZQ

	ChannelID := "UCfpcfju9rBs5o_xQLXmLQHQ"
	tm := internal.NewTaskPool()

	channels := []string{ChannelID}
	platforms := []platform.Name{platform.YOUTUBE}

	subscriberId := "pewdiepie_exe"
	subPlatform := platform.INSTAGRAM
	err = tm.AddTask(channels, platforms, subscriberId, subPlatform, *database.NewUserRepository(db, subscriberId, subPlatform))
	if err != nil {
		return
	}

	tm.RunAll()
}
