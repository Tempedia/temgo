package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
	"gitlab.com/wiky.lyu/temgo/x"
)

type FindTemtemLocationsRequest struct {
	Query    string `json:"query" form:"query" query:"query"`
	Page     int    `json:"page" form:"page" query:"page"`
	PageSize int    `json:"pageSize" form:"pageSize" query:"pageSize"`
}

func FindTemtemLocations(c echo.Context) error {
	ctx := c.(*middleware.Context)

	req := FindTemtemLocationsRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest()
	}

	req.Page, req.PageSize = x.Pagination(req.Page, req.PageSize)

	locations, total, err := temtemdb.FindTemtemLocations(req.Query, req.Page, req.PageSize)
	if err != nil {
		return err
	}
	return ctx.List(locations, req.Page, req.PageSize, total)
}

func FindTemtemLocationAreasByLocation(c echo.Context) error {
	ctx := c.(*middleware.Context)
	location := ctx.Param(`name`)

	areas, err := temtemdb.FindTemtemLocationAreasByLocation(location)
	if err != nil {
		return err
	}
	return ctx.Success(areas)
}

/* 查询temtem所在的区域 */
func FindTemtemLocationsByTemtem(c echo.Context) error {
	ctx := c.(*middleware.Context)
	name, _ := parseTemtemName(ctx.Param(`name`))

	temtem, err := temtemdb.GetTemtemByName(name)
	if err != nil {
		return err
	} else if temtem == nil {
		return ctx.NotFound()
	}
	locations, err := temtemdb.FindTemtemLocationsByTemtem(temtem.Name)
	if err != nil {
		return err
	}
	return ctx.Success(locations)
}
