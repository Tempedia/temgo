package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
)

func GetTemtemTrait(c echo.Context) error {
	ctx := c.(*middleware.Context)

	name := ctx.Param(`name`)

	trait, err := temtemdb.GetTemtemTrait(name)
	if err != nil {
		return err
	}
	return ctx.SuccessOr404(trait)
}
