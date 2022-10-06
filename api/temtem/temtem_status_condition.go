package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
	"gitlab.com/wiky.lyu/temgo/x"
)

type FindTemtemStatusConditionsRequest struct {
	Query    string `json:"query" form:"query" query:"query"`
	Group    string `json:"group" form:"group" query:"group"`
	Page     int    `json:"page" form:"page" query:"page"`
	PageSize int    `json:"pageSize" form:"pageSize" query:"pageSize"`
}

func FindTemtemStatusConditions(c echo.Context) error {
	ctx := c.(*middleware.Context)

	req := FindTemtemStatusConditionsRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest()
	}

	req.Page, req.PageSize = x.Pagination(req.Page, req.PageSize)

	conditions, total, err := temtemdb.FindTemtemStatusConditions(req.Query, req.Group, req.Page, req.PageSize)
	if err != nil {
		return err
	}
	return ctx.List(conditions, req.Page, req.PageSize, total)
}

func FindTemtemStatusConditionTechniques(c echo.Context) error {
	ctx := c.(*middleware.Context)
	name := ctx.Param(`name`)

	techniques, err := temtemdb.FindTemtemTechniquesByStatusCondition(name)
	if err != nil {
		return err
	}
	return ctx.Success(techniques)
}

func FindTemtemStatusConditionTraits(c echo.Context) error {
	ctx := c.(*middleware.Context)
	name := ctx.Param(`name`)

	traits, err := temtemdb.FindTemtemTraitsByStatusCondition(name)
	if err != nil {
		return err
	}
	return ctx.Success(traits)
}
