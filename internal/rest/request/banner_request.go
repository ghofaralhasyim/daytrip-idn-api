package rest_request

type InsertBannerRequest struct {
	DesktopImage string `form:"desktop_image" validate:"required"`
	MobileImage  string `form:"mobile_image" validate:"required"`
	Cta          string `form:"cta"`
	CtaUrl       string `form:"cta_url"`
	Title        string `form:"title"`
	Description  string `form:"description"`
}
