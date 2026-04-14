package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"taskflow/internal/config"
	"taskflow/internal/handlers"
	"taskflow/internal/middleware"
	"taskflow/internal/repository"
	"taskflow/internal/service"
)

func main() {
	cfg := config.LoadConfig()
	db := config.NewDBPool(cfg.DBUrl)
	defer db.Close()

	// Initialize Auth layers
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	// authHandler := handlers.NewAuthHandler(userService)
	authHandler := handlers.NewAuthHandler(userService, cfg.JWTSecret)

	// Initialize Project layers
	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo)
	projectHandler := handlers.NewProjectHandler(projectService)

	// Initialize Task layers
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	r := gin.Default()

	// Public routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))

	protected.PATCH("/tasks/:id", taskHandler.UpdateTask)
	protected.DELETE("/tasks/:id", taskHandler.DeleteTask)

	protected.GET("/projects/:id", projectHandler.GetProjectByID)
	protected.PATCH("/projects/:id", projectHandler.UpdateProject)
	protected.DELETE("/projects/:id", projectHandler.DeleteProject)

	protected.GET("/protected", func(c *gin.Context) {
		userID, _ := c.Get("user_id")

		c.JSON(200, gin.H{
			"message": "You are authenticated",
			"user_id": userID,
		})
	})

	// Project routes
	protected.POST("/projects", projectHandler.CreateProject)
	protected.GET("/projects", projectHandler.GetProjects)

	// Task routes
	protected.POST("/projects/:id/tasks", taskHandler.CreateTask)
	protected.GET("/projects/:id/tasks", taskHandler.GetTasks)

	log.Println("Server running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
