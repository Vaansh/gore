package handler

import (
	"github.com/Vaansh/gore/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskHandler struct {
	TaskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
	}
}

func (th *TaskHandler) RunTask(c *gin.Context) {
	channels := c.Query("channels")
	subscriberID := c.Query("subscriberId")
	igPostTags := c.Query("igPostTags")

	err := th.TaskService.RunTask(channels, subscriberID, igPostTags)
	if err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}

	c.String(http.StatusOK, "Task started successfully.")
}

func (th *TaskHandler) StopTask(c *gin.Context) {
	subscriberID := c.Param("subscriberId")
	subscriberPlatform := c.Param("subscriberPlatform")

	err := th.TaskService.StopTask(subscriberID, subscriberPlatform)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.String(http.StatusOK, "Task stopped successfully.")
}
