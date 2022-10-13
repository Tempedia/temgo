package google

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/wiky.lyu/temgo/api/middleware"
	"gitlab.com/wiky.lyu/temgo/service/google"
)

type ValidateIAPRequst struct {
	ProductID string `json:"product_id" form:"product_id" query:"product_id" validate:"required"`
	Token     string `json:"token" form:"token" query:"token" validate:"required"`
}

/* 验证IAP购买 */
func ValidateIAP(c echo.Context) error {
	ctx := c.(*middleware.Context)
	req := ValidateIAPRequst{}
	if err := ctx.BindAndValidate(&req); err != nil {
		return ctx.BadRequest()
	}

	purchased, err := google.ValidateProduct(req.ProductID, req.Token)
	if err != nil {
		return err
	}
	return ctx.Success(purchased)
}
