package rest_request

type InvitationResponseRequest struct {
	InvitationId int64  `json:"invitation_id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	IsAttending  string `json:"is_attending" validate:"required"`
	Message      string `json:"message,omitempty"`
}
