package files

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/service/files"
)

func Register(router *echo.Group) {

	router.Static("/", files.Folder())
}
