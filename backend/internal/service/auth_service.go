// internal/service/auth_service.go
package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"jimsyyap/ndb/internal/config"
	"jimsyyap/ndb/internal/domain"
	"jimsyyap/ndb/internal/repository"
)

type AuthService struct {
	userRepo domain.UserRepository
	cfg      *config.Config
}

func NewAuthService(userRepo domain.UserRepository, cfg *config.Config) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

// Claims represents JWT claims
type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Create creates a new user with hashed password
func (s *AuthService) Create(user *domain.User, password string) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return s.userRepo.Create(user)
}

// GetByID retrieves a user by ID
func (s *AuthService) GetByID(id uuid.UUID) (*domain.User, error) {
	return s.userRepo.GetByID(id)
}

// GetByEmail retrieves a user by email
func (s *AuthService) GetByEmail(email string) (*domain.User, error) {
	return s.userRepo.GetByEmail(email)
}

// Update updates a user
func (s *AuthService) Update(user *domain.User) error {
	return s.userRepo.Update(user)
}

// Delete deletes a user
func (s *AuthService) Delete(id uuid.UUID) error {
	return s.userRepo.Delete(id)
}

// List lists users with pagination
func (s *AuthService) List(limit, offset int) ([]*domain.User, error) {
	return s.userRepo.List(limit, offset)
}

// Authenticate authenticates a user and returns a JWT token
func (s *AuthService) Authenticate(email, password string) (*domain.User, string, error) {
	// Get user by email
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := s.generateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

// ValidateToken validates a JWT token and returns the user
func (s *AuthService) ValidateToken(tokenString string) (*domain.User, error) {
	// Parse token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Get user by ID
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID in token")
	}

	return s.userRepo.GetByID(userID)
}

// generateToken generates a JWT token for a user
func (s *AuthService) generateToken(user *domain.User) (string, error) {
	// Set expiration time
	expirationTime := time.Now().Add(time.Duration(s.cfg.JWT.ExpirationHours) * time.Hour)

	// Create claims
	claims := &Claims{
		UserID: user.ID.String(),
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tennis-club-api",
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
