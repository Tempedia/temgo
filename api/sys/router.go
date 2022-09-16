package sys

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
)

func Register(router *echo.Group) {
	router.GET("/status", GetSystemStatus, middleware.TryStaffAuthMiddleware)
	router.POST("/superstaff", CreateSuperStaff)
}
