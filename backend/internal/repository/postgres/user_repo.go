// internal/repository/postgres/user_repo.go
package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"jimsyyap/nbh/internal/domain"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *domain.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, name, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	// Generate UUID if not provided
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	_, err := r.db.Exec(
		query,
		user.ID,
		user.Email,
		user.Password,
		user.Name,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

func (r *UserRepository) GetByID(id uuid.UUID) (*domain.User, error) {
	query := `
		SELECT id, email, password_hash, name, role, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user domain.User
	err := r.db.Get(&user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	query := `
		SELECT id, email, password_hash as password, name, role, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	var user domain.User
	err := r.db.Get(&user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *domain.User) error {
	query := `
		UPDATE users
		SET email = $1, password_hash = $2, name = $3, role = $4, updated_at = $5
		WHERE id = $6
	`

	user.UpdatedAt = time.Now()

	_, err := r.db.Exec(
		query,
		user.Email,
		user.Password,
		user.Name,
		user.Role,
		user.UpdatedAt,
		user.ID,
	)

	return err
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepository) List(limit, offset int) ([]*domain.User, error) {
	query := `
		SELECT id, email, name, role, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	var users []*domain.User
	err := r.db.Select(&users, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return users, nil
}
