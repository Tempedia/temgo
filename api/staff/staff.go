package staff

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
)

func GetSelf(c echo.Context) error {
	ctx := c.(*middleware.Context)
	return ctx.Success(ctx.Staff)
}
