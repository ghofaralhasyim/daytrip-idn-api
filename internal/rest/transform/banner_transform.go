package response_transform

import (
	"github.com/daytrip-idn-api/internal/entities"
)

type BannerResponse struct {
	Id           int64  `json:"id"`
	DesktopImage string `json:"desktopImage"`
	MobileImage  string `json:"mobileImage"`
	Cta          string `json:"cta"`
	CtaUrl       string `json:"ctaUrl"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

func TransformBannerResponse(
	e *entities.BannerEntity,
) *BannerResponse {
	return &BannerResponse{
		Id:           e.Id,
		DesktopImage: e.DesktopImage,
		MobileImage:  e.MobileImage,
		Cta:          e.Cta,
		CtaUrl:       e.CtaUrl,
		Title:        e.Title,
		Description:  e.Description,
	}
}

func TransformListBannerResponse(
	data []entities.BannerEntity,
) []BannerResponse {
	result := make([]BannerResponse, 0)
	for _, item := range data {
		transform := TransformBannerResponse(&item)
		result = append(result, *transform)
	}

	return result
}
