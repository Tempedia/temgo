package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
	"gitlab.com/wiky.lyu/temgo/x"
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

type FindTemtemTraitsRequest struct {
	Query    string `json:"query" form:"query" query:"query"`
	Page     int    `json:"page" form:"page" query:"page"`
	PageSize int    `json:"pageSize" form:"pageSize" query:"pageSize"`
}

func FindTemtemTraits(c echo.Context) error {
	ctx := c.(*middleware.Context)
	req := FindTemtemTraitsRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest()
	}
	req.Page, req.PageSize = x.Pagination(req.Page, req.PageSize)

	traits, total, err := temtemdb.FindTemtemTraits(req.Query, req.Page, req.PageSize)
	if err != nil {
		return err
	}
	return ctx.List(traits, req.Page, req.PageSize, total)
}

func FindTemtemsByTrait(c echo.Context) error {
	ctx := c.(*middleware.Context)

	name := ctx.Param(`name`)

	// time.Sleep(time.Second * 3)

	temtems, err := temtemdb.FindTemtemsByTrait(name)
	if err != nil {
		return err
	}
	return ctx.Success(temtems)
}
