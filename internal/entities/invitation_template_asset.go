package entities

import "time"

type InvitationTemplateAssetEntity struct {
	Id         int64
	TemplateId int64
	AssetType  string
	AssetUrl   string
	SortOrder  int
	CreatedAt  time.Time
}

func MakeInvitationTemplateAssetEntity(
	id int64,
	templateId int64,
	assetType string,
	assetUrl string,
	sortOrder int,
	createdAt time.Time,
) *InvitationTemplateAssetEntity {
	return &InvitationTemplateAssetEntity{
		Id:         id,
		TemplateId: templateId,
		AssetType:  assetType,
		AssetUrl:   assetUrl,
		SortOrder:  sortOrder,
		CreatedAt:  createdAt,
	}
}
