package dysmsapi

import (
	"micro-shop-api/goods-web/initialize"
	"micro-shop-api/goods-web/utils"
	"testing"
)

func TestSendRegisterSms(t *testing.T) {
	initialize.InitConfig()
	err := SendRegisterSms("18169630262", utils.GenValidateCode(6))
	if err != nil {
		panic(err)
	}
}
