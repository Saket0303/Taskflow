package handlers

import (
	"net/http"

	"taskflow/internal/models"
	"taskflow/internal/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// Create Task
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var t models.Task

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed"})
		return
	}

	projectID := c.Param("id")
	t.ProjectID = projectID

	err := h.service.CreateTask(c.Request.Context(), &t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, t)
}

// Get Tasks
func (h *TaskHandler) GetTasks(c *gin.Context) {
	projectID := c.Param("id")
	status := c.Query("status")
	assignee := c.Query("assignee")

	tasks, err := h.service.GetTasks(c.Request.Context(), projectID, status, assignee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	// var t models.Task

	var t models.UpdateTaskRequest

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed"})
		return
	}

	err := h.service.UpdateTask(c.Request.Context(), id, &t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task updated"})
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	userID, _ := c.Get("user_id")

	// Get task
	_, err := h.service.GetTaskByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	// Check project ownership
	// (Simplified: assuming project owner = current user)
	// You can improve later

	if userID.(string) == "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	err = h.service.DeleteTask(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete task"})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
	c.Status(http.StatusNoContent)
}
