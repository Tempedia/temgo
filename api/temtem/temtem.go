package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
	"gitlab.com/wiky.lyu/temgo/x"
)

type FindTemtemsRequest struct {
	Query    string   `json:"query" form:"query" query:"query"`
	Type     []string `json:"type" form:"type" query:"type"`
	Sort     string   `json:"sort" form:"sort" query:"sort"`
	Page     int      `json:"page" form:"page" query:"page"`
	PageSize int      `json:"pageSize" form:"pageSize" query:"pageSize"`
}

func FindTemtems(c echo.Context) error {
	ctx := c.(*middleware.Context)
	req := FindTemtemsRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest()
	}

	req.Page, req.PageSize = x.Pagination(req.Page, req.PageSize)

	temtems, total, err := temtemdb.FindTemtems(req.Query, req.Type, req.Sort, req.Page, req.PageSize)
	if err != nil {
		return err
	}
	return ctx.List(temtems, req.Page, req.PageSize, total)
}

/* 查询进化前的Temtem */
func FindTemtemsEvolvesFrom(c echo.Context) error {
	ctx := c.(*middleware.Context)
	name := ctx.Param(`name`)

	temtems, err := temtemdb.FindTemtemsEvolvesFrom(name)
	if err != nil {
		return err
	}
	return ctx.Success(temtems)
}

func GetTemtem(c echo.Context) error {
	ctx := c.(*middleware.Context)
	name := ctx.Param(`name`)
	temtem, err := temtemdb.GetTemtemByName(name)
	if err != nil {
		return err
	}
	return ctx.SuccessOr404(temtem)
}
