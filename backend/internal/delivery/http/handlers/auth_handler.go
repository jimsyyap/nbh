// internal/delivery/http/handlers/auth_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"jimsyyap/nbh/internal/domain"
	"jimsyyap/nbh/internal/delivery/dto"
)

type AuthHandler struct {
	userService domain.UserService
}

func NewAuthHandler(userService domain.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if email already exists
	_, err := h.userService.GetByEmail(req.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	// Create user
	user := &domain.User{
		Email: req.Email,
		Name:  req.Name,
		Role:  "member", // Default role for new users
	}

	err = h.userService.Create(user, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Authenticate user
	user, token, err := h.userService.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Return user info and token
	c.JSON(http.StatusOK, gin.H{
		"user": dto.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
			Role:  user.Role,
		},
		"token": token,
	})
}

// GetMe retrieves the current user
func (h *AuthHandler) GetMe(c *gin.Context) {
	// Get user from context (set by auth middleware)
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	u := user.(*domain.User)
	c.JSON(http.StatusOK, dto.UserResponse{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
		Role:  u.Role,
	})
}
