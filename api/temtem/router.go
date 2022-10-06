package temtem

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
)

func Register(router *echo.Group) {
	authRouter := router.Group("", middleware.StaffAuthMiddleware)

	/* 属性 */
	router.GET("/types", FindTemtemTypes)
	router.GET("/type/:name", GetTemtemType)
	router.GET("/type/:name/icon", GetTemtemTypeIcon)
	authRouter.PUT("/type/:name", UpdateTemtemType)

	/* Temtem */
	router.GET("/temtems", FindTemtems)
	router.GET("/temtem/:name", GetTemtem)
	router.GET("/temtem/:name/evolves_from", FindTemtemsEvolvesFrom)
	router.GET("/temtem/:name/locations", FindTemtemLocationsByTemtem)

	/* Trait */
	router.GET("/traits", FindTemtemTraits)
	router.GET("/trait/:name", GetTemtemTrait)
	router.GET("/trait/:name/temtems", FindTemtemsByTrait)

	/* Technique */
	router.GET("/temtem/:name/techniques", FindTemtemTechniquesByTemtem)
	router.GET("/techniques", FindTemtemTechniques)
	router.GET("/technique/:name/temtems", FindTemtemsByTechnique)

	/* Location */
	router.GET("/locations", FindTemtemLocations)
	router.GET("/location/:name/areas", FindTemtemLocationAreasByLocation)

	/* Condition */
	router.GET("/status/conditions", FindTemtemStatusConditions)
	router.GET("/status/condition/:name/techniques", FindTemtemStatusConditionTechniques)
	router.GET("/status/condition/:name/traits", FindTemtemStatusConditionTraits)
}
