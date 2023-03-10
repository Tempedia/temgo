package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/files"
	"gitlab.com/wiky.lyu/temgo/api/google"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	"gitlab.com/wiky.lyu/temgo/api/temtem"
	"gitlab.com/wiky.lyu/temgo/api/user"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Register(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
	router := e.Group("/api", middleware.APIMiddleware)
	temtem.Register(router.Group("/temtem"))
	files.Register(router.Group("/files"))
	google.Register(router.Group("/google"))
	user.Register(router.Group("/user"))
}
