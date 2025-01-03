package repository

import (
	"database/sql"

	"github.com/QDEX-Core/oneart-identity-service/internal/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *domain.User) error {
	query := `INSERT INTO users (email, password, created_at, updated_at)
              VALUES ($1, $2, NOW(), NOW())
              RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, user.Email, user.Password).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}

func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	query := `SELECT id, email, password, created_at, updated_at
              FROM users
              WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
