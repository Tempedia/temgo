package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
)

func Register(router *echo.Group) {
	authRouter := router.Group("", middleware.StaffAuthMiddleware)

	router.GET("/types", FindTemtemTypes)
	router.GET("/type/:name", GetTemtemType)

	authRouter.PUT("/type/:name", UpdateTemtemType)
}
