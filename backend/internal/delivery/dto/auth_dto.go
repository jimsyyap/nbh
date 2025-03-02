// internal/delivery/dto/auth_dto.go
package dto

import (
	"github.com/google/uuid"
)

// RegisterRequest represents the request for user registration
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required"`
}

// LoginRequest represents the request for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserResponse represents the response for user data
type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
	Role  string    `json:"role"`
}

// internal/config/config.go
package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config represents the application configuration
type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
	}
	JWT struct {
		Secret         string
		ExpirationHours int
	}
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	cfg := &Config{}

	// Server configuration
	cfg.Server.Port = getEnv("SERVER_PORT", "8080")

	// Database configuration
	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "5432")
	cfg.Database.User = getEnv("DB_USER", "postgres")
	cfg.Database.Password = getEnv("DB_PASSWORD", "postgres")
	cfg.Database.Name = getEnv("DB_NAME", "tennis_club")
	cfg.Database.SSLMode = getEnv("DB_SSL_MODE", "disable")

	// JWT configuration
	cfg.JWT.Secret = getEnv("JWT_SECRET", "your-secret-key")
	cfg.JWT.ExpirationHours = getEnvAsInt("JWT_EXPIRATION_HOURS", 24)

	return cfg, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvAsInt gets an environment variable as an integer or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// migrations/000001_create_users_table.up.sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  role VARCHAR(50) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create indexes
CREATE INDEX idx_users_email ON users(email);

-- Create initial admin user (password: admin123)
INSERT INTO users (id, email, password_hash, name, role, created_at, updated_at)
VALUES (
  uuid_generate_v4(),
  'admin@tennisclub.com',
  '$2a$10$1G4DMQ2Td44JC4hCj2eI9.oO7zx/HGdGJuqgvk3qXyYFx1N6hBpym',
  'Admin User',
  'admin',
  NOW(),
  NOW()
);
