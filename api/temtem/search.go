package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
)

type SearchRequest struct {
	Query string `json:"query" form:"query" query:"query"`
}

func Search(c echo.Context) error {
	ctx := c.(*middleware.Context)

	req := SearchRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest()
	}

	r := make(map[string]interface{})
	if req.Query == "" {
		return ctx.Success(r)
	}
	pageSize := 3

	temtems, _, err := temtemdb.FindTemtems(req.Query, nil, "", "", 1, pageSize)
	if err != nil {
		return err
	}
	r["temtems"] = temtems

	traits, _, err := temtemdb.FindTemtemTraits(req.Query, 1, pageSize)
	if err != nil {
		return err
	}
	r["traits"] = traits

	techniques, _, err := temtemdb.FindTemtemTechniques(req.Query, nil, "", 1, pageSize)
	if err != nil {
		return err
	}
	r["techniques"] = techniques

	locations, _, err := temtemdb.FindTemtemLocations(req.Query, 1, pageSize)
	if err != nil {
		return err
	}
	r["locations"] = locations

	conditions, _, err := temtemdb.FindTemtemStatusConditions(req.Query, "", 1, pageSize)
	if err != nil {
		return err
	}
	r["conditions"] = conditions

	return ctx.Success(r)
}
