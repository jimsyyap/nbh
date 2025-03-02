// internal/delivery/http/middleware/auth_middleware.go
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"jimsyyap/nbh/internal/domain"
)

type AuthMiddleware struct {
	userService domain.UserService
}

func NewAuthMiddleware(userService domain.UserService) *AuthMiddleware {
	return &AuthMiddleware{
		userService: userService,
	}
}

// Authenticate middleware checks if the request has a valid JWT token
func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Check if the header has the Bearer scheme
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Extract token
		tokenString := headerParts[1]

		// Validate token
		user, err := m.userService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user in context
		c.Set("user", user)
		c.Next()
	}
}

// RequireRole middleware checks if the user has the required role
func (m *AuthMiddleware) RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user from context
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		// Check role
		if user.(*domain.User).Role != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
