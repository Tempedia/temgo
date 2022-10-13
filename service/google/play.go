package google

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/awa/go-iap/playstore"
)

var _client *playstore.Client
var _packageName string

func Init(packageName, jsonFile string) error {
	jsonKey, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	_client, err = playstore.New(jsonKey)
	if err != nil {
		return err
	}
	_packageName = packageName
	return nil
}

/* 验证IAP购买 */
func ValidateProduct(productID, token string) (bool, error) {
	resp, err := _client.VerifyProduct(context.Background(), _packageName, productID, token)
	if err != nil {
		return false, err
	} else if resp == nil {
		return false, fmt.Errorf("未知错误")
	}
	return resp.PurchaseState == 0, nil
}
