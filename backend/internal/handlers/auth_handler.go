package handlers

import (
	"net/http"

	"taskflow/internal/models"
	"taskflow/internal/service"
	"taskflow/internal/utils"

	"github.com/gin-gonic/gin"
)

// type AuthHandler struct {
// 	service *service.UserService
// }

type AuthHandler struct {
	service *service.UserService
	secret  string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func NewAuthHandler(service *service.UserService) *AuthHandler {
// 	return &AuthHandler{service: service}
// }

func NewAuthHandler(service *service.UserService, secret string) *AuthHandler {
	return &AuthHandler{
		service: service,
		secret:  secret,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "validation failed",
		})
		return
	}

	err := h.service.Register(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "validation failed",
		})
		return
	}

	user, err := h.service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// token, err := utils.GenerateToken(user.ID, user.Email, "supersecret")
	token, err := utils.GenerateToken(user.ID, user.Email, h.secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
