package api

import (
	"github.com/Vaansh/gore"
	"github.com/Vaansh/gore/internal/domain"
	"github.com/Vaansh/gore/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskHandler struct {
	TaskService *domain.TaskService
}

func NewTaskHandler(taskService *domain.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
	}
}

func (th *TaskHandler) RunInstagramTask(c *gin.Context) {
	var request RunInstagramTaskRequest
	if err := c.ShouldBindJSON(&request); err != nil || len(request.PublisherIds) != len(request.Sources) {
		c.JSON(http.StatusBadRequest, TaskResponse{Success: false, Error: err.Error()})
		return
	}

	err := th.TaskService.RunTask(request.PublisherIds, request.Sources, request.SubscriberId, go_pubsub.INSTAGRAM,
		*model.NewInstagramMetaData(request.IgUserId, request.LongLivedAccessToken, request.Hashtags, request.Frequency))
	if err != nil {
		c.JSON(http.StatusBadRequest, TaskResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, TaskResponse{Success: true})
}

func (th *TaskHandler) StopTask(c *gin.Context) {
	subscriberId := c.Param("id")
	subscriberPlatform := c.Param("platform")

	err := th.TaskService.StopTask(subscriberId, subscriberPlatform)
	if err != nil {
		c.JSON(http.StatusNotFound, TaskResponse{Success: false, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, TaskResponse{Success: true})
}
