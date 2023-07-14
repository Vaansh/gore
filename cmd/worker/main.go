package main

import (
	"fmt"
	"github.com/Vaansh/gore/internal/handler"
	"github.com/Vaansh/gore/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//channels := []string{ChannelID}
	//platforms := []platform.Name{platform.YOUTUBE}

	//ChannelID := "UCQ4zIVlfhsmvds7WuKeL2Bw"
	//subscriberId := "pewdiepie_exe"
	//hashtags := "#pewdiepie #memes #dankmemes #meme #funny #memesdaily #dank #funnymemes #pewdiepiememes #memereview #pewds #spicymemes #dankmeme #lwiay #reels #reelsinstagram #instagram #explore #viral #trending #tiktok #shorts #youtube #fyp #gamer #dailymemes #gaming #mrbeast"

	//subPlatform := platform.INSTAGRAM
	//task := domain.NewTask(channels, platforms, subscriberId, subPlatform, model.MetaData{IgPostTags: hashtags}, *database.NewUserRepository(db, subscriberId, subPlatform))

	//stop := make(chan struct{})
	//go task.Run(stop)
	//time.Sleep(3 * time.Second)
	//stop <- struct{}{}

	taskService, err := service.NewTaskService()
	if err != nil {
		log.Fatalf("Error initializing task service: %s", err)
	}

	taskHandler := handler.NewTaskHandler(taskService)

	// Set up Gin router
	router := gin.Default()

	// Define endpoints and their corresponding handlers
	router.POST("/tasks", taskHandler.RunTask)
	router.DELETE("/tasks/:subscriberPlatform/:subscriberId", taskHandler.StopTask)

	// Start the HTTP server
	serverPort := "8080"
	fmt.Printf("Server listening on port %s\n", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, router))
}
