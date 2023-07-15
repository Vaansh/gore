package main

import (
	"fmt"
	"github.com/Vaansh/gore/internal/api"
	"github.com/Vaansh/gore/internal/domain"
	"github.com/Vaansh/gore/internal/gcloud"
	"github.com/Vaansh/gore/internal/util"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := gcloud.InitLogger(); err != nil {
		gcloud.LogFatal(fmt.Sprintf("Failed to initialize logger: %v", err))
	}

	if err := gcloud.InitStorage(); err != nil {
		gcloud.LogFatal(fmt.Sprintf("Failed to initialize storage: %v", err))
	}

	// create instance of service and handler
	taskService := domain.NewTaskService()
	taskHandler := api.NewTaskHandler(taskService)

	// server config
	router := gin.Default()
	serverPort := util.Getenv("SERVER_PORT", false)
	if serverPort == "" {
		serverPort = "8080"
	}

	// register routes
	router.Use(api.AuthMiddleware)
	router.POST("/tasks/ig", taskHandler.RunInstagramTask)
	router.DELETE("/tasks/:platform/:id", taskHandler.StopTask)

	gcloud.LogInfo(fmt.Sprintf("Server listening on port %s\n", serverPort))
	err := router.Run(":" + serverPort)
	if err != nil {
		gcloud.LogFatal(err.Error())
	}
}
