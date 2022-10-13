package google

import (
	"github.com/labstack/echo/v4"
)

func Register(router *echo.Group) {
	router.GET("/play/iap/validate", ValidateIAP)
}
