package staff

import (
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/errors"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	staffdb "gitlab.com/wiky.lyu/temgo/db/staff"
)

type LoginRequest struct {
	Username string `json:"username" query:"username" form:"username" validate:"required"`
	Password string `json:"password" query:"password" form:"password" validate:"required"`
}

func Login(c echo.Context) error {
	ctx := c.(*middleware.Context)
	req := LoginRequest{}
	if err := ctx.BindAndValidate(&req); err != nil {
		return ctx.BadRequest()
	}

	staff, err := staffdb.GetStaffByUsername(req.Username)
	if err != nil {
		return ctx.InternalServerError()
	} else if staff == nil { /* 用户不存在 */
		return ctx.Fail(errors.ErrStaffUsernameNotExists, nil)
	} else if staff.Status == staffdb.StaffStatusBanned {
		return ctx.Fail(errors.ErrStaffStatusBanned, nil)
	}

	if !staff.Auth(req.Password) {
		return ctx.Fail(errors.ErrStaffPasswordIncorrect, nil)
	}

	expiresAt := time.Now().AddDate(0, 0, 1)
	token, err := staffdb.CreateStaffToken(staff.ID, expiresAt, ctx.Request().UserAgent(), ctx.RealIP())
	if err != nil {
		return ctx.InternalServerError()
	}

	sess, _ := session.Get(middleware.StaffCookie, ctx)
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 0,
	}

	sess.Values[middleware.StaffTokenKey] = token.ID
	sess.Save(ctx.Request(), ctx.Response())

	return ctx.Success(nil)
}

func Logout(c echo.Context) error {
	ctx := c.(*middleware.Context)

	sess, _ := session.Get(middleware.StaffCookie, ctx)
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 0,
	}

	sess.Values[middleware.StaffTokenKey] = ""
	sess.Save(ctx.Request(), ctx.Response())

	return ctx.Success(nil)
}
