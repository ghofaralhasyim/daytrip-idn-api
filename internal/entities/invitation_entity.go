package entities

import (
	"mime/multipart"
	"time"
)

type InvitationEntity struct {
	Id          int64
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
	Image       *string
	Image1      *string
	KeyPass     string
	BirthdayVal *int64

	// req image
	ImageFile  *multipart.FileHeader
	ImageFile1 *multipart.FileHeader
}

func MakeInvitationEntity(
	id int64,
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
	image, image1 *string,
	keyPass string,
	birthdayVal *int64,
) InvitationEntity {
	return InvitationEntity{
		Id:          id,
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
		Image:       image,
		Image1:      image1,
		KeyPass:     keyPass,
		BirthdayVal: birthdayVal,
	}
}
