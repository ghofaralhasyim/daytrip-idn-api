package routes

import (
	"github.com/daytrip-idn-api/internal/middleware"
	"github.com/daytrip-idn-api/internal/modules"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, m *modules.AppModules) {
	apiv1 := e.Group("/v1")
	apiv1.Static("/images", "public/images")

	apiv1.POST("/login", m.Controllers.User.Login)

	// ========================== BANNERS ROUTES ======================== //

	apiv1.GET("/banners", m.Controllers.Banner.GetBanners)
	apiv1.POST("/banners", m.Controllers.Banner.CreateBanners, middleware.AuthMiddleware)
	apiv1.PUT("/banners", m.Controllers.Banner.UpdateBanner, middleware.AuthMiddleware)
	apiv1.DELETE("/banners/:id", m.Controllers.Banner.DeleteBanner, middleware.AuthMiddleware)

	apiv1.GET("/destinations", m.Controllers.Destination.GetDestinations)

	// ========================== MESSAGES ROUTES ======================== //

	apiv1.GET("/messages",
		m.Controllers.Message.GetMessages,
		middleware.AuthMiddleware,
	)
	apiv1.POST("/messages",
		m.Controllers.Message.InsertMessage,
	)
	apiv1.DELETE("/message/:messageId",
		m.Controllers.Message.DeleteMessage,
		middleware.AuthMiddleware,
	)

	apiv1.GET("/activity",
		m.Controllers.Activity.GetActivities,
		middleware.AuthMiddleware,
	)

	// ========================== INVITATION ROUTES ======================== //
	apiv1.GET("/invitations",
		m.Controllers.Invitation.GetInvitations, middleware.AuthMiddleware,
	)

	apiv1.POST("/invitations",
		m.Controllers.Invitation.CreateInvitation, middleware.AuthMiddleware,
	)

	apiv1.DELETE("/invitations/:id",
		m.Controllers.Invitation.DeleteInvitation, middleware.AuthMiddleware,
	)

	apiv1.GET("/invitations/:slug",
		m.Controllers.Invitation.GetInvitationBySlug,
	)

	apiv1.GET("/admin/invitations/:slug",
		m.Controllers.Invitation.GetAdminInvitationBySlug, middleware.AuthMiddleware,
	)

	apiv1.PUT("/invitations/:slug",
		m.Controllers.Invitation.UpdateInvitation, middleware.AuthMiddleware,
	)

	apiv1.POST("/invitations/attendance",
		m.Controllers.Invitation.InsertAttendance, middleware.AuthMiddleware,
	)

	apiv1.GET("/invitations/attendance/:slug",
		m.Controllers.Invitation.GetAttendanceBySlug,
	)
}
