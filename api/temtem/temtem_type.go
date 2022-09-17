package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
	"gitlab.com/wiky.lyu/temgo/x"
)

func FindTemtemTypes(c echo.Context) error {
	ctx := c.(*middleware.Context)

	types, err := temtemdb.FindTemtemTypes()
	if err != nil {
		return ctx.InternalServerError()
	}
	return ctx.Success(types)
}

func GetTemtemType(c echo.Context) error {
	ctx := c.(*middleware.Context)

	name := ctx.Param(`name`)

	t, err := temtemdb.GetTemtemType(name)
	if err != nil {
		return err
	}
	return ctx.SuccessOr404(t)
}

type UpdateTemtemTypeRequest struct {
	Comment string   `json:"comment" form:"comment" query:"comment"`
	Trivia  []string `json:"trivia" form:"trivia" query:"trivia"`
}

func UpdateTemtemType(c echo.Context) error {
	ctx := c.(*middleware.Context)
	req := UpdateTemtemTypeRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest()
	}

	name := ctx.Param(`name`)

	req.Trivia = x.StripStringArray(req.Trivia)

	t, err := temtemdb.UpdateTemtemType(name, req.Comment, req.Trivia)
	if err != nil {
		return err
	}
	return ctx.Success(t)
}
