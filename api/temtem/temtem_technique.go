package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	temtemdb "gitlab.com/wiky.lyu/temgo/db/temtem"
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
