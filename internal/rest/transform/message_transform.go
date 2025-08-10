package response_transform

import (
	"time"

	"github.com/daytrip-idn-api/internal/entities"
)

type MessageResponse struct {
	Id          int64     `json:"id"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	PackageName string    `json:"packageName"`
	Message     string    `json:"message"`
	CreatedAt   time.Time `json:"createdAt"`
}

func TransformMessageResponse(
	e *entities.MessageEntity,
) *MessageResponse {
	return &MessageResponse{
		Id:          e.Id,
		Phone:       e.Phone,
		Email:       e.Email,
		PackageName: e.PackageName,
		Message:     e.Message,
		CreatedAt:   e.CreatedAt,
	}
}

func TransformListMessageResponse(
	data []entities.MessageEntity,
) []MessageResponse {
	result := make([]MessageResponse, 0)
	for _, item := range data {
		transform := TransformMessageResponse(&item)
		result = append(result, *transform)
	}

	return result
}
