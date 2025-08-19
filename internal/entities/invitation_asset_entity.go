package entities

import (
	"mime/multipart"
	"time"
)

type InvitationAssetEntity struct {
	Id           int64
	InvitationId int64
	AssetUrl     string
	SortOrder    int
	CreatedAt    time.Time
	FileHeader   multipart.FileHeader
}

func MakeInvitationAssetEntity(
	id int64,
	invitationId int64,
	assetUrl string,
	sortOrder int,
	createdAt time.Time,
) InvitationAssetEntity {
	return InvitationAssetEntity{
		Id:           id,
		InvitationId: invitationId,
		AssetUrl:     assetUrl,
		SortOrder:    sortOrder,
		CreatedAt:    createdAt,
	}
}
