package google

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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
		log.Errorf("购买验证失败:%v", err)
		return err
	}
	return ctx.Success(purchased)
}
