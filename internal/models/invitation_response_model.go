package models

import (
	"time"

	"github.com/daytrip-idn-api/internal/entities"
)

type InvitationResponse struct {
	Id           int64     `db:"id"`
	InvitationId int64     `db:"invitation_id"`
	Name         string    `db:"name"`
	IsAttending  string    `db:"is_attending"`
	Message      string    `db:"message"`
	CreatedAt    time.Time `db:"created_at"`
}

func ToInvitationResponseModel(e entities.InvitationResponseEntity) InvitationResponse {
	return InvitationResponse{
		Id:           e.Id,
		InvitationId: e.InvitationId,
		Name:         e.Name,
		IsAttending:  e.IsAttending,
		Message:      e.Message,
		CreatedAt:    e.CreatedAt,
	}
}
