package main

import (
	"github.com/Vaansh/gore/internal/domain"
	"github.com/Vaansh/gore/internal/rest"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// create instance of the task service
	taskService, err := domain.NewTaskService()
	if err != nil {
		log.Fatalf("Error initializing task service: %s", err)
	}

	// server config
	serverPort := "8080"
	router := gin.Default()
	taskHandler := rest.NewTaskHandler(taskService)

	// register routes
	router.POST("/tasks/ig", taskHandler.RunInstagramTask)
	router.DELETE("/tasks/:platform/:id", taskHandler.StopTask)

	log.Printf("Server listening on port %s\n", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, router))
}
