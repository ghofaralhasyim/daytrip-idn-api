package rest_request

import "time"

type InvitationResponseRequest struct {
	InvitationId int64  `json:"invitation_id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	IsAttending  string `json:"is_attending" validate:"required"`
	Message      string `json:"message,omitempty"`
}

type InvtationRequest struct {
	Title       string    `form:"title" validate:"required"`
	Description string    `form:"description" validate:"required"`
	TemplateId  int64     `form:"templateId" validate:"required"`
	StartDate   time.Time `form:"startDate" validate:"required"`
	EndDate     time.Time `form:"endDate" validate:"required"`
	MapsUrl     string    `form:"mapsUrl" validate:"required"`
	Address     string    `form:"address" validate:"required"`
	Location    string    `form:"location" validate:"required"`
	DressCode   string    `form:"dressCode" validate:"required"`
	KeyPass     string    `form:"keyPass" validate:"required"`
	BirthdayVal int64     `form:"birthdayVal"`
}
