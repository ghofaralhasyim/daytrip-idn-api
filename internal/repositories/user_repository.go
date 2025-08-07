package repositories

import (
	"database/sql"
	"time"

	"github.com/daytrip-idn-api/internal/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(userId int) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT 
			u.user_id, u.name, u.email, u.password_hash, u.created_at, u.updated_at
		FROM
			users u
		WHERE 
			u.email = $1;
	`

	var user models.User
	var updated sql.NullString

	err := r.db.QueryRow(query, email).Scan(&user.UserId, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &updated)
	if err != nil {
		return nil, err
	}

	if updated.Valid {
		layout := "2006-01-02T15:04:05Z"

		parsedTime, err := time.Parse(layout, updated.String)
		if err != nil {
			return nil, err
		}

		user.UpdatedAt = parsedTime
	}

	return &user, nil
}

func (r *userRepository) GetUserById(userId int) (*models.User, error) {
	query := `
		SELECT 
			u.user_id, u.name, u.email, u.password_hash, u.created_at, u.updated_at
		FROM
			users u
		WHERE 
			u.user_id = $1;
	`

	var user models.User
	var updated sql.NullString

	err := r.db.QueryRow(query, user).Scan(&user.UserId, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &updated)
	if err != nil {
		return nil, err
	}

	if updated.Valid {
		layout := "2006-01-02T15:04:05Z"

		parsedTime, err := time.Parse(layout, updated.String)
		if err != nil {
			return nil, err
		}

		user.UpdatedAt = parsedTime
	}

	return &user, nil
}
