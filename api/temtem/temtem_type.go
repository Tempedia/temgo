package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
)

func FindTemtemTypes(c echo.Context) error {
	ctx := c.(*middleware.Context)

	types, err := temtemdb.FindTemtemTypes()
	if err != nil {
		return ctx.InternalServerError()
	}
	return ctx.Success(types)
}
