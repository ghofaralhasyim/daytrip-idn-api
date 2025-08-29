package models

import (
	"time"

	"github.com/daytrip-idn-api/internal/entities"
)

type Message struct {
	Id          int64     `db:"id"`
	Phone       string    `db:"phone"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	PackageName string    `db:"package_name"`
	Message     string    `db:"message"`
	CreatedAt   time.Time `db:"created_at"`
}

func ToMessageModel(
	e entities.MessageEntity,
) Message {
	return Message{
		Id:          e.Id,
		Phone:       e.Phone,
		Name:        e.Name,
		Email:       e.Email,
		PackageName: e.PackageName,
		Message:     e.Message,
		CreatedAt:   e.CreatedAt,
	}
}
