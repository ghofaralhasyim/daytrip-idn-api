package rest_request

type InsertBannerRequest struct {
	DesktopImage string `form:"desktopImage" validate:"required"`
	MobileImage  string `form:"mobileImage" validate:"required"`
	Cta          string `form:"cta"`
	CtaUrl       string `form:"ctaUrl"`
	Title        string `form:"title"`
	Description  string `form:"description"`
}

type UpdateBannerRequest struct {
	Id           int64  `form:"id" validate:"required"`
	DesktopImage string `form:"desktopImage"`
	MobileImage  string `form:"mobileImage"`
	Cta          string `form:"cta"`
	CtaUrl       string `form:"ctaUrl"`
	Title        string `form:"title"`
	Description  string `form:"description"`
}
