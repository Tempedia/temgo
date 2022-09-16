package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	staffdb "gitlab.com/wiky.lyu/temgo/db/staff"
)

const (
	StaffCookie   = "StaffCookie"
	StaffTokenKey = "Token"
)

func StaffAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*Context)
		tokenID := ctx.GetSession(StaffCookie, StaffTokenKey)
		if tokenID == "" {
			return ctx.Unauthorized()
		}
		token, err := staffdb.GetStaffToken(tokenID)
		if err != nil {
			return ctx.InternalServerError()
		} else if token == nil {
			return ctx.Unauthorized()
		}

		/* 检查token有效性 */
		if token.Status != staffdb.StaffTokenStatusOK || (!token.ExpiresAt.IsZero() && token.ExpiresAt.Before(time.Now())) {
			return ctx.Unauthorized()
		}

		/* 检查用户有效性 */
		if token.Staff == nil || token.Staff.Status == staffdb.StaffStatusBanned {
			return ctx.Forbidden()
		}

		ctx.Staff = token.Staff

		return next(ctx)
	}
}

/* 尝试解析登录 */
func TryStaffAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*Context)
		tokenID := ctx.GetSession(StaffCookie, StaffTokenKey)
		if tokenID == "" {
			return next(ctx)
		}
		token, err := staffdb.GetStaffToken(tokenID)
		if err != nil {
			return next(ctx)
		} else if token == nil {
			return next(ctx)
		}

		/* 检查token有效性 */
		if token.Status != staffdb.StaffTokenStatusOK || (!token.ExpiresAt.IsZero() && token.ExpiresAt.Before(time.Now())) {
			return next(ctx)
		}

		/* 检查用户有效性 */
		if token.Staff == nil || token.Staff.Status == staffdb.StaffStatusBanned {
			return next(ctx)
		}

		ctx.Staff = token.Staff

		return next(ctx)
	}
}
