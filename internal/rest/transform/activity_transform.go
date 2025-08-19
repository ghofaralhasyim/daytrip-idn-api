package response_transform

import "github.com/daytrip-idn-api/internal/entities"

type ActivityResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func TransformActivityResponse(
	e *entities.ActivityEntity,
) *ActivityResponse {
	return &ActivityResponse{
		Id:    e.Id,
		Name:  e.Name,
		Image: e.Image,
	}
}

func TransformListActivityResponse(
	data []entities.ActivityEntity,
) []ActivityResponse {
	result := make([]ActivityResponse, 0)
	for _, item := range data {
		transform := TransformActivityResponse(&item)
		result = append(result, *transform)
	}

	return result
}
