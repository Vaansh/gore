package main

import (
	"github.com/Vaansh/gore/internal/api"
	"github.com/Vaansh/gore/internal/domain"
	"github.com/Vaansh/gore/internal/util"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
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

	log.Printf("Server listening on port %s\n", serverPort)
	err := router.Run(":" + serverPort)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
