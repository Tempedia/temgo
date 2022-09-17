package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/files"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	"gitlab.com/wiky.lyu/temgo/api/staff"
	"gitlab.com/wiky.lyu/temgo/api/sys"
	"gitlab.com/wiky.lyu/temgo/api/temtem"
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
	sys.Register(router.Group("/sys"))
	staff.Register(router.Group("/staff"))
	temtem.Register(router.Group("/temtem"))
	files.Register(router.Group("/files"))
}
