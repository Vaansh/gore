package main

import (
	"fmt"
	"log"

	"github.com/Vaansh/gore/internal/api"
	"github.com/Vaansh/gore/internal/domain"
	"github.com/Vaansh/gore/internal/gcloud"
	"github.com/Vaansh/gore/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading environment variables file: missing or incorrect .env file")
	}

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
	port := util.Getenv("PORT", false)
	if port == "" {
		port = "8080"
	}

	// register routes
	router.Use(api.AuthMiddleware)
	router.POST("/tasks/ig", taskHandler.RunInstagramTask)
	router.DELETE("/tasks/:platform/:id", taskHandler.StopTask)

	host, err := util.GetLocalIP()
	if err != nil {
		gcloud.LogWarning(err.Error())
	}

	gcloud.LogInfo(fmt.Sprintf("Server listening on: http://%s:%s\n", host, port))

	err = router.Run(":" + port)
	if err != nil {
		gcloud.LogFatal(err.Error())
	}
}
