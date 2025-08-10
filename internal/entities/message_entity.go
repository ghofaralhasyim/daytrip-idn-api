package entities

import "time"

type MessageEntity struct {
	Id          int64
	Phone       string
	Email       string
	PackageName string
	Message     string
	CreatedAt   time.Time
}

func MakeMessageEntity(
	Id int64,
	Phone string,
	Email string,
	PackageName string,
	Message string,
	CreatedAt time.Time,
) *MessageEntity {
	return &MessageEntity{
		Id:          Id,
		Phone:       Phone,
		Email:       Email,
		PackageName: PackageName,
		Message:     Message,
		CreatedAt:   CreatedAt,
	}
}
