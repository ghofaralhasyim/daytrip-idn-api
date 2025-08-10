package models

import "github.com/daytrip-idn-api/internal/entities"

type Banner struct {
	Id           int64  `db:"id"`
	DesktopImage string `db:"desktop_image"`
	MobileImage  string `db:"mobile_image"`
	Cta          string `db:"cta"`
	CtaUrl       string `db:"cta_url"`
	Title        string `db:"title"`
	Description  string `db:"description"`
}

func ToBannerModel(
	e entities.BannerEntity,
) Banner {
	return Banner{
		Id:           e.Id,
		DesktopImage: e.DesktopImage,
		MobileImage:  e.MobileImage,
		Cta:          e.Cta,
		CtaUrl:       e.CtaUrl,
		Title:        e.Title,
		Description:  e.Description,
	}
}
