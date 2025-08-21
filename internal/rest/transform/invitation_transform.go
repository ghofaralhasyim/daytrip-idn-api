package response_transform

import (
	"time"

	"github.com/daytrip-idn-api/internal/entities"
)

type InvitationEntityResponse struct {
	Id          int64      `json:"id"`
	Slug        string     `json:"slug"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	TemplateId  *int64     `json:"templateId"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
	MapsUrl     *string    `json:"mapsUrl"`
	Address     *string    `json:"address"`
	Location    *string    `json:"location"`
	DressCode   *string    `json:"dressCode"`
	CreatedAt   time.Time  `json:"createdAt"`
	Image       *string    `json:"image"`
	Image1      *string    `json:"image1"`
	KeyPass     string     `json:"keyPass,omitempty"`
}

func TransformInvitationResponse(
	e *entities.InvitationEntity,
) *InvitationEntityResponse {
	return &InvitationEntityResponse{
		Id:          e.Id,
		Slug:        e.Slug,
		Title:       e.Title,
		Description: e.Description,
		TemplateId:  e.TemplateId,
		StartDate:   e.StartDate,
		EndDate:     e.EndDate,
		MapsUrl:     e.MapsUrl,
		Address:     e.Address,
		Location:    e.Location,
		DressCode:   e.DressCode,
		CreatedAt:   e.CreatedAt,
		Image:       e.Image,
		Image1:      e.Image1,
		KeyPass:     e.KeyPass,
	}
}

func TransformListInvitationResponse(
	data []entities.InvitationEntity,
) []InvitationEntityResponse {
	result := make([]InvitationEntityResponse, 0)
	for _, item := range data {
		transform := TransformInvitationResponse(&item)
		result = append(result, *transform)
	}

	return result
}
