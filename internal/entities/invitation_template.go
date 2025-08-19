package entities

import "time"

type InvitationTemplateEntity struct {
	Id                    int64
	Name                  string
	Description           *string
	InvitationAssetsCount int
	CreatedAt             time.Time
}

func MakeInvitationTemplateEntity(
	id int64,
	name string,
	description *string,
	assetsCount int,
	createdAt time.Time,
) *InvitationTemplateEntity {
	return &InvitationTemplateEntity{
		Id:                    id,
		Name:                  name,
		Description:           description,
		InvitationAssetsCount: assetsCount,
		CreatedAt:             createdAt,
	}
}
