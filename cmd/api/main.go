package main

import (
	"fmt"
	"github.com/Vaansh/gore/internal/api"
	"github.com/Vaansh/gore/internal/domain"
	"github.com/Vaansh/gore/internal/gcloud"
	"github.com/Vaansh/gore/internal/util"
	"github.com/gin-gonic/gin"
	"net"
	"os"
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
	port := util.Getenv("PORT", false)
	if port == "" {
		port = "8080"
	}

	// register routes
	router.Use(api.AuthMiddleware)
	router.POST("/tasks/ig", taskHandler.RunInstagramTask)
	router.DELETE("/tasks/:platform/:id", taskHandler.StopTask)

	fmt.Println(getLocalIP())

	host := getLocalIP()
	gcloud.LogInfo(fmt.Sprintf("Server listening on: http://%s:%s\n", host, port))

	err := router.Run(":" + port)
	if err != nil {
		gcloud.LogFatal(err.Error())
	}
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error retrieving the local IP address:", err)
		os.Exit(1)
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}

	fmt.Println("Unable to find the local IP address.")
	return ""
}
