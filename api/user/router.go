package user

import (
	"github.com/labstack/echo/v4"
)

func Register(router *echo.Group) {
	router.POST("/team", CreateTemtemUserTeam)
	router.GET("/team/:id", GetTemtemUserTeam)
}
