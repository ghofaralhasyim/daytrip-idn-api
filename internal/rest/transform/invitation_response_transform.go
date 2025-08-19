package response_transform

import (
	"time"

	"github.com/daytrip-idn-api/internal/entities"
)

type InvitationResponseResponse struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Message     string    `json:"message"`
	IsAttending string    `json:"isAttending"`
	CreatedAt   time.Time `json:"createdAt"`
}

func TransformInvitationResponseResponse(
	e *entities.InvitationResponseEntity,
) *InvitationResponseResponse {
	return &InvitationResponseResponse{
		Id:          e.Id,
		Name:        e.Name,
		Message:     e.Message,
		IsAttending: e.IsAttending,
		CreatedAt:   e.CreatedAt,
	}
}

func TransformListInvitationResponseResponse(
	data []entities.InvitationResponseEntity,
) []InvitationResponseResponse {
	result := make([]InvitationResponseResponse, 0)
	for _, item := range data {
		transform := TransformInvitationResponseResponse(&item)
		result = append(result, *transform)
	}

	return result
}
