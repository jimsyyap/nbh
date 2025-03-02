// internal/domain/user.go
package domain

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Password is never returned in JSON
	Name      string    `json:"name"`
	Role      string    `json:"role"` // "admin", "member", "guest"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository defines the methods that any user repository must implement
type UserRepository interface {
	Create(user *User) error
	GetByID(id uuid.UUID) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uuid.UUID) error
	List(limit, offset int) ([]*User, error)
}

// UserService defines the methods that the user service must implement
type UserService interface {
	Create(user *User, password string) error
	GetByID(id uuid.UUID) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uuid.UUID) error
	List(limit, offset int) ([]*User, error)
	Authenticate(email, password string) (*User, string, error) // Returns user, token, error
	ValidateToken(token string) (*User, error)
}

// Token represents a JWT token in the system
type Token struct {
	UserID    uuid.UUID `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

