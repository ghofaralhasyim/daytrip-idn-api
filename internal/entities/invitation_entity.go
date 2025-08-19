package entities

import "time"

type InvitationEntity struct {
	Id          int64
	UserId      int64
	Slug        string
	Title       string
	Description *string
	TemplateId  *int64
	StartDate   *time.Time
	EndDate     *time.Time
	MapsUrl     *string
	Address     *string
	Location    *string
	DressCode   *string
	CreatedAt   time.Time
	Assets      []InvitationAssetEntity
}

func MakeInvitationEntity(
	id int64,
	userId int64,
	slug string,
	title string,
	description *string,
	templateId *int64,
	startDate *time.Time,
	endDate *time.Time,
	mapsUrl *string,
	address *string,
	location *string,
	dressCode *string,
	createdAt time.Time,
) InvitationEntity {
	return InvitationEntity{
		Id:          id,
		UserId:      userId,
		Slug:        slug,
		Title:       title,
		Description: description,
		TemplateId:  templateId,
		StartDate:   startDate,
		EndDate:     endDate,
		MapsUrl:     mapsUrl,
		Address:     address,
		Location:    location,
		DressCode:   dressCode,
		CreatedAt:   createdAt,
	}
}
