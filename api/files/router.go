package files

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	"gitlab.com/wiky.lyu/temgo/service/files"
)

func Register(router *echo.Group) {
	authRouter := router.Group("", middleware.StaffAuthMiddleware)

	authRouter.POST("/file", Upload)
	router.Static("/", files.Folder())
}
