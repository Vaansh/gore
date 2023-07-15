package main

import (
	"github.com/Vaansh/gore/internal/api"
	"github.com/Vaansh/gore/internal/domain"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// create instance of the task api
	taskService, err := domain.NewTaskService()
	if err != nil {
		log.Fatalf("Error initializing task api: %s", err)
	}
	taskHandler := api.NewTaskHandler(taskService)

	// server config
	serverPort := ":8080"
	router := gin.Default()

	// register routes
	router.Use(api.AuthMiddleware)
	router.POST("/tasks/ig", taskHandler.RunInstagramTask)
	router.DELETE("/tasks/:platform/:id", taskHandler.StopTask)

	log.Printf("Server listening on port %s\n", serverPort)
	err = router.Run(serverPort)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
