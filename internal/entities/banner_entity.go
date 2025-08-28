package entities

import "mime/multipart"

type BannerEntity struct {
	Id           int64
	DesktopImage string
	MobileImage  string
	Cta          string
	CtaUrl       string
	Title        string
	Description  string

	//
	MobileImageFile  *multipart.FileHeader
	DesktopImageFile *multipart.FileHeader
}

func MakeBannerEntity(
	id int64,
	desktopImage string,
	mobileImage string,
	cta string,
	ctaUrl string,
	title string,
	description string,
) *BannerEntity {
	return &BannerEntity{
		Id:           id,
		DesktopImage: desktopImage,
		MobileImage:  mobileImage,
		Cta:          cta,
		CtaUrl:       ctaUrl,
		Title:        title,
		Description:  description,
	}
}
