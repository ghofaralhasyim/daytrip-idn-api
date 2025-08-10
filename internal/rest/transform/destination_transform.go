package response_transform

import (
	"time"

	"github.com/daytrip-idn-api/internal/entities"
)

type DestinationResponse struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
}

func TransformDestinationResponse(
	e *entities.DestinationEntity,
) *DestinationResponse {
	return &DestinationResponse{
		Id:        e.Id,
		Name:      e.Name,
		Image:     e.Image,
		CreatedAt: e.CreatedAt,
	}
}

func TransformListDestinationResponse(
	data []entities.DestinationEntity,
) []DestinationResponse {
	result := make([]DestinationResponse, 0)
	for _, item := range data {
		transform := TransformDestinationResponse(&item)
		result = append(result, *transform)
	}

	return result
}
