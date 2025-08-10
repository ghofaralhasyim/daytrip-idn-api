package repositories

import (
	"context"
	"database/sql"

	"github.com/daytrip-idn-api/internal/models"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT 
			u.id, u.name, u.email, u.phone,
			u.image, u.password_hash, u.role, 
			u.created_at
		FROM
			users u
		WHERE 
			u.email = $1;
	`

	var user models.User
	var image sql.NullString

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.Id, &user.Name, &user.Email,
		&user.Phone, &image, &user.PasswordHash,
		&user.Role, &user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	if image.Valid {
		user.Image = image.String
	}

	return &user, nil
}
