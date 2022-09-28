package middleware

import (
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	staffdb "gitlab.com/wiky.lyu/temgo/db/staff"
)

type Context struct {
	echo.Context

	Staff *staffdb.Staff
}

func (c *Context) Param(key string) string {
	v := c.Context.Param(key)
	r, err := url.PathUnescape(v)
	if err != nil {
		return v
	}
	return r
}

func (c *Context) IntParam(key string) int64 {
	v := c.Param(key)

	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

func (c *Context) IntFormParam(key string) int64 {
	v := c.FormValue(key)
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

func (c *Context) Success(data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 0,
		"data":   data,
	})
}

func (c *Context) SuccessOr404(data interface{}) error {
	if data != nil {
		return c.Success(data)
	} else {
		return c.NotFound()
	}
}

func (c *Context) Fail(status int, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
		"data":   data,
	})
}

func (c *Context) List(data interface{}, page, pageSize, total int, v ...interface{}) error {
	d := map[string]interface{}{
		"page":      page,
		"page_size": pageSize,
		"total":     total,
		"list":      data,
	}
	if len(v) > 0 {
		d["extra"] = v[0]
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 0,
		"data":   d,
	})
}

func (c *Context) Bind(data interface{}) error {
	if err := c.Context.Bind(data); err != nil {
		return err
	}

	/* 过滤字符串的首尾空格 */
	v := reflect.ValueOf(data).Elem()
	for i := 0; i < v.NumField(); i++ {
		vv := v.Field(i)
		if vv.Kind() == reflect.String {
			vv.SetString(strings.TrimSpace(vv.String()))
		}
	}
	return nil
}

func (c *Context) BindAndValidate(data interface{}) error {
	if err := c.Bind(data); err != nil {
		return err
	}

	return c.Validate(data)
}

func (c *Context) BadRequest() error {
	return c.NoContent(http.StatusBadRequest)
}

func (c *Context) NotFound() error {
	return c.NoContent(http.StatusNotFound)
}

func (c *Context) InternalServerError() error {
	return c.NoContent(http.StatusInternalServerError)
}

func (c *Context) Unauthorized() error {
	return c.NoContent(http.StatusUnauthorized)
}

func (c *Context) Forbidden() error {
	return c.NoContent(http.StatusForbidden)
}

func (ctx *Context) GetSession(name, value string) string {
	sess, _ := session.Get(name, ctx)
	if sess == nil {
		return ""
	}
	v, ok := sess.Values[value]
	if !ok || v == nil {
		return ""
	}
	return v.(string)
}

// func (c *Context) XLSX(filename string, wb *xlsx.File) error {

// 	buf := bytes.NewBuffer(nil)
// 	if err := wb.Write(buf); err != nil {
// 		return err
// 	}

// 	c.Response().Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s.xlsx", filename))
// 	return c.Stream(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf)
// }
