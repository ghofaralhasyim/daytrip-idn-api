package entities

import "time"

type InvitationResponseEntity struct {
	Id           int64
	InvitationId int64
	Name         string
	IsAttending  string
	Message      string
	CreatedAt    time.Time
}

func MakeInvitationResponseEntity(
	id int64,
	invitationId int64,
	name string,
	isAttending string,
	message string,
	createdAt time.Time,
) *InvitationResponseEntity {
	return &InvitationResponseEntity{
		Id:           id,
		InvitationId: invitationId,
		Name:         name,
		IsAttending:  isAttending,
		Message:      message,
		CreatedAt:    createdAt,
	}
}
