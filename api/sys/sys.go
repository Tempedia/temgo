package sys

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/errors"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	staffdb "gitlab.com/wiky.lyu/temgo/db/staff"
)

/*
 * 获取系统状态
 */
func GetSystemStatus(c echo.Context) error {
	ctx := c.(*middleware.Context)

	staff, err := staffdb.GetSuperStaff()
	if err != nil {
		return err
	}

	return ctx.Success(map[string]interface{}{
		"has_superuser": staff != nil,
		"login":         ctx.Staff != nil,
	})
}

type CreateSuperuserRequest struct {
	Username string `json:"username" form:"username" query:"username" validate:"gte=6"`
	Password string `json:"password" form:"password" query:"password" validate:"gte=6"`

	Name  string `json:"name" form:"name" query:"name" validate:"gt=0"`
	Phone string `json:"phone" form:"phone" query:"phone"`
	Email string `json:"email" form:"email" query:"email"`
}

/* 创建超级用户，只有不存在超级用户时才能创建 */
func CreateSuperStaff(c echo.Context) error {
	ctx := c.(*middleware.Context)
	req := CreateSuperuserRequest{}
	if err := ctx.BindAndValidate(&req); err != nil {
		fmt.Printf("%v\n", err)
		return ctx.BadRequest()
	}

	staff, err := staffdb.GetStaffByUsername(req.Username)
	if err != nil {
		return err
	} else if staff != nil {
		/* 用户名已存在 */
		return ctx.Fail(errors.ErrStaffUsernameExists, nil)
	}

	staff, err = staffdb.CreateSuperStaff(req.Username, req.Password, req.Name, req.Phone, req.Email)
	if err != nil {
		return err
	}
	return ctx.Success(staff)
}
