package middleware

import (
	"github.com/labstack/echo/v4"
)

func APIMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := Context{Context: c}
		/* 检查IP地址是否有效 */
		// if white, err := xModel.GetIPWhiteList(ctx.RealIP()); err != nil {
		// 	return ctx.InternalServerError()
		// } else if !white {
		// 	return ctx.Forbidden()
		// }
		return next(&ctx)
	}
}
