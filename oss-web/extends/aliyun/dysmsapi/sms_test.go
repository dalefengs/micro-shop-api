package dysmsapi

import (
	"micro-shop-api/oss-web/initialize"
	"micro-shop-api/oss-web/utils"
	"testing"
)

func TestSendRegisterSms(t *testing.T) {
	initialize.InitConfig()
	err := SendRegisterSms("18169630262", utils.GenValidateCode(6))
	if err != nil {
		panic(err)
	}
}
