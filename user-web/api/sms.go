package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"micro-shop-api/user-web/global"
	"micro-shop-api/user-web/utils"
	redisutils "micro-shop-api/user-web/utils/redis"
	"time"

	"micro-shop-api/user-web/extends/aliyun/dysmsapi"
	"micro-shop-api/user-web/global/response"
	status2 "micro-shop-api/user-web/global/status"
	"micro-shop-api/user-web/validator/forms"
)

// SendSmsCode 发送短信
func SendSmsCode(c *gin.Context) {
	SendCodeForm := forms.SendCodeForm{}
	err := c.ShouldBindJSON(&SendCodeForm)
	if err != nil {
		if err == io.EOF {
			response.Fail(c, status2.InvalidParameter)
			return
		}
		fmt.Println(err)
		ve := err.(validator.ValidationErrors)
		response.FailCodeMsg(c, status2.InvalidParameter, ve.Error())
		return
	}
	code := utils.GenValidateCode(6)
	switch SendCodeForm.CodeType {
	case 1: // 注册
		if utils.IsDev() {
			code = "123456"
		} else {
			if err = dysmsapi.SendRegisterSms(SendCodeForm.Mobile, code); err != nil {
				response.FailCodeMsg(c, status2.SmsSendCodeFail, err.Error())
				return
			}
		}

	//case 2: // 找回密码
	default:
		response.Fail(c, status2.InvalidParameter)
		return
	}
	// 保存短信到Redis中
	r := redisutils.NewRedisCline()
	// 写入Redis 有效期为5秒
	err = r.Set(context.Background(), global.Config.Redis.Prefix.Register+SendCodeForm.Mobile, code,
		time.Duration(global.Config.Aliyun.Sms.Expire)*time.Second).Err()
	if err != nil {
		response.FailCodeMsg(c, status2.SmsSendCodeFail, err.Error())
		return
	}

	response.SuccessMsg(c, "发送短信成功")
}
