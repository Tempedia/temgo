package temtem

import (
	"github.com/labstack/echo/v4"
)

func Register(router *echo.Group) {
	// authRouter := router.Group("", middleware.StaffAuthMiddleware)

	router.GET("/types", FindTemtemTypes)
	router.GET("/type/:name", GetTemtemType)
}
