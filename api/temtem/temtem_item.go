package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
	"gitlab.com/wiky.lyu/temgo/x"
)

type FindTemtemItemsRequest struct {
	Query    string   `json:"query" form:"query" query:"query"`
	Category []string `json:"category" form:"category" query:"category"`
	Page     int      `json:"page" form:"page" query:"page"`
	PageSize int      `json:"pageSize" form:"pageSize" query:"pageSize"`
}

func FindTemtemItems(c echo.Context) error {
	ctx := c.(*middleware.Context)

	req := FindTemtemItemsRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest()
	}
	req.Page, req.PageSize = x.Pagination(req.Page, req.PageSize)

	items, total, err := temtemdb.FindTemtemItems(req.Query, req.Category, req.Page, req.PageSize)
	if err != nil {
		return err
	}
	return ctx.List(items, req.Page, req.PageSize, total)
}
