package user

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
	"gitlab.com/wiky.lyu/temgo/service/host"
)

type CreateUserTeamRequest struct {
	Name    string                   `json:"name" form:"name" query:"name"`
	Temtems []map[string]interface{} `json:"temtems" form:"temtems" query:"temtems" validate:"gt=0"`
}

func CreateTemtemUserTeam(c echo.Context) error {
	ctx := c.(*middleware.Context)
	req := CreateUserTeamRequest{}
	if err := ctx.BindAndValidate(&req); err != nil {
		return ctx.BadRequest()
	}
	team, err := temtemdb.CreateUserTeam(req.Name, req.Temtems)
	if err != nil {
		return err
	}
	return ctx.Success(map[string]interface{}{
		"share_url": host.URL(fmt.Sprintf("/shared/team/%s", team.ID)),
	})
}

func GetTemtemUserTeam(c echo.Context) error {
	ctx := c.(*middleware.Context)
	id := ctx.Param(`id`)

	team, err := temtemdb.GetTemtemUserTeam(id)
	if err != nil {
		return err
	} else if team == nil {
		return ctx.NotFound()
	}
	return ctx.Success(team)
}
