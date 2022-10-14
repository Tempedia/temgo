package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
	"gitlab.com/wiky.lyu/temgo/x"
)

/* 获取temtem的所有技能 */
func FindTemtemTechniquesByTemtem(c echo.Context) error {
	ctx := c.(*middleware.Context)

	name := ctx.Param(`name`)

	levelingUpTechniques, err := temtemdb.FindTemtemLevelingUpTechniques(name)
	if err != nil {
		return err
	}

	courseTechniques, err := temtemdb.FindTemtemCourseTechniques(name)
	if err != nil {
		return err
	}

	breedingTechniques, err := temtemdb.FindTemtemBreedingTechniques(name)
	if err != nil {
		return err
	}

	return ctx.Success(map[string]interface{}{
		"leveling_up": levelingUpTechniques,
		"course":      courseTechniques,
		"breeding":    breedingTechniques,
	})
}

type FindTemtemTechniquesRequest struct {
	Query    string   `json:"query" form:"query" query:"query"`
	Type     []string `json:"type" form:"type" query:"type"`
	Class    string   `json:"class" form:"class" query:"class"`
	Page     int      `json:"page" form:"page" query:"page"`
	PageSize int      `json:"pageSize" form:"pageSize" query:"pageSize"`
}

func FindTemtemTechniques(c echo.Context) error {
	ctx := c.(*middleware.Context)

	req := FindTemtemTechniquesRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest()
	}

	req.Page, req.PageSize = x.Pagination(req.Page, req.PageSize)

	techniques, total, err := temtemdb.FindTemtemTechniques(req.Query, req.Type, req.Class, req.Page, req.PageSize)
	if err != nil {
		return err
	}

	return ctx.List(techniques, req.Page, req.PageSize, total)
}

func FindTemtemsByTechnique(c echo.Context) error {
	ctx := c.(*middleware.Context)
	name := ctx.Param(`name`)

	levelingUp, err := temtemdb.FindTemtemsByLevelingUpTechnique(name)
	if err != nil {
		return err
	}
	course, err := temtemdb.FindTemtemsByCourseTechnique(name)
	if err != nil {
		return err
	}
	breeding, err := temtemdb.FindTemtemsByBreedingTechnique(name)
	if err != nil {
		return err
	}

	return ctx.Success(map[string]interface{}{
		"leveling_up": levelingUp,
		"course":      course,
		"breeding":    breeding,
	})
}

type FindTemtemCourseItemsRequest struct {
	Query    string `json:"query" form:"query" query:"query"`
	Page     int    `json:"page" form:"page" query:"page"`
	PageSize int    `json:"pageSize" form:"pageSize" query:"pageSize"`
}

/* 获取技能教程 */
func FindTemtemCourseItems(c echo.Context) error {
	ctx := c.(*middleware.Context)
	req := FindTemtemCourseItemsRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest()
	}
	req.Page, req.PageSize = x.Pagination(req.Page, req.PageSize)

	items, total, err := temtemdb.FindTemtemCourseItems(req.Query, req.Page, req.PageSize)
	if err != nil {
		return err
	}
	return ctx.List(items, req.Page, req.PageSize, total)
}
