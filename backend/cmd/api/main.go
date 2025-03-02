// cmd/api/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"jimsyyap/nbh/internal/config"
	"jimsyyap/nbh/internal/delivery/http/handlers"
	"jimsyyap/nbh/internal/delivery/http/middleware"
	"jimsyyap/nbh/internal/repository/postgres"
	"jimsyyap/nbh/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, cfg)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Initialize router
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		
		c.Next()
	})

	// Define routes
	// Public routes
	router.POST("/api/auth/register", authHandler.Register)
	router.POST("/api/auth/login", authHandler.Login)

	// Protected routes
	authorized := router.Group("/api")
	authorized.Use(authMiddleware.Authenticate())
	{
		authorized.GET("/auth/me", authHandler.GetMe)

		// Admin only routes
		admin := authorized.Group("/admin")
		admin.Use(authMiddleware.RequireRole("admin"))
		{
			// Add admin routes here
		}
	}

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server started on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
