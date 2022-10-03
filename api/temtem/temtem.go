package temtem

import (
	"strings"

	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
	"gitlab.com/wiky.lyu/temgo/x"
)

type FindTemtemsRequest struct {
	Query    string   `json:"query" form:"query" query:"query"`
	Type     []string `json:"type" form:"type" query:"type"`
	Trait    string   `json:"trait" form:"trait" query:"trait"`
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

	temtems, total, err := temtemdb.FindTemtems(req.Query, req.Type, req.Trait, req.Sort, req.Page, req.PageSize)
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

	subspecie := ""
	seps := strings.Split(name, "#")
	if len(seps) > 1 {
		name = seps[0]
		subspecie = seps[1]
	} else {
		seps = strings.Split(name, "(")
		if len(seps) > 1 {
			name = strings.TrimSpace(seps[0])
			subspecie = strings.Trim(seps[1], ")")
		}
	}

	temtem, err := temtemdb.GetTemtemByName(name)
	if err != nil {
		return err
	}
	if temtem != nil && subspecie != "" {
		for _, sub := range temtem.Subspecies {
			if sub.Type == subspecie {
				temtem.Icon = sub.Icon
				temtem.LumaIcon = sub.LumaIcon
				break
			}
		}
	}
	return ctx.SuccessOr404(temtem)
}
