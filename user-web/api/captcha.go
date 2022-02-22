package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"micro-shop-api/user-web/global/response"
	status2 "micro-shop-api/user-web/global/status"
)

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	var driver base64Captcha.Driver
	driver = base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64, err := captcha.Generate()
	if err != nil {
		response.FailCodeMsg(c, status2.Fail, err.Error())
		return
	}
	data := make(map[string]string)
	data["id"] = id
	data["img"] = b64
	fmt.Println(b64)
	response.SuccessData(c, data)
}

func CaptchaVerify(id, code string) bool {
	return store.Verify(id, code, false)
}
