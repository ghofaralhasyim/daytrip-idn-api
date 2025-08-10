package models

import "time"

type User struct {
	Id           int64
	Name         string
	Email        string
	Phone        string
	Image        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
}
