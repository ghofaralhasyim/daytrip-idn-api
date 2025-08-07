package models

import "time"

type User struct {
	UserId       int       `json:"user_id"`
	Name         string    `json:"name"`
	Email        string    `json:"email" validate:"email"`
	PasswordHash string    `json:"passowrd_hash,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
