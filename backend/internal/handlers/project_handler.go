package handlers

import (
	"net/http"

	"taskflow/internal/models"
	"taskflow/internal/service"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	service *service.ProjectService
}

type UpdateProjectRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func NewProjectHandler(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: service}
}

// Create Project
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var p models.Project

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed"})
		return
	}

	userID, _ := c.Get("user_id")
	p.OwnerID = userID.(string)

	err := h.service.CreateProject(c.Request.Context(), &p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, p)
}

// Get Projects
func (h *ProjectHandler) GetProjects(c *gin.Context) {
	userID, _ := c.Get("user_id")

	projects, err := h.service.GetProjects(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

func (h *ProjectHandler) GetProjectByID(c *gin.Context) {
	id := c.Param("id")

	project, tasks, err := h.service.GetProjectByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          project.ID,
		"name":        project.Name,
		"description": project.Description,
		"owner_id":    project.OwnerID,
		"tasks":       tasks,
	})
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	id := c.Param("id")

	var req UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation failed"})
		return
	}

	err := h.service.UpdateProject(c.Request.Context(), id, req.Name, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "project updated"})
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteProject(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete project"})
		return
	}

	// c.JSON(http.StatusOK, gin.H{"message": "project deleted"})
	c.Status(http.StatusNoContent)
}
