package dysmsapi

import (
	"context"
	"errors"
	"fmt"
	redisutils "micro-shop-api/order-web/utils/redis"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"go.uber.org/zap"

	"micro-shop-api/order-web/global"
)

// CreateClient
/* 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient() (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &global.Config.Aliyun.AccessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: &global.Config.Aliyun.AccessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

// SendRegisterSms 发送用户注册短信
func SendRegisterSms(mobile, code string) error {
	client, err := CreateClient()
	if err != nil {
		return err
	}
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers: tea.String(mobile),
		SignName:     tea.String(global.Config.Aliyun.Sms.SignName),
		TemplateCode: tea.String(global.Config.Aliyun.Sms.TemplateCode),
		//TemplateParam: tea.String("{\"code\":\"123456\"}"),
		TemplateParam: tea.String(fmt.Sprintf(`{"code": "%s"}`, code)),
	}
	// 复制代码运行请自行打印 API 的返回值
	result, err := client.SendSms(sendSmsRequest)
	if err != nil {
		return err
	}
	if tea.StringValue(result.Body.Code) != "OK" {
		zap.S().Info(result, err)
		return errors.New(tea.StringValue(result.Body.Message))
	}
	return nil
}

// VerifySmsCode 验证短信验证码
// _type 1注册，2找回密码
func VerifySmsCode(_type int, mobile, code string) bool {
	var prefix string
	switch _type {
	case 1:
		prefix = global.Config.Redis.Prefix.Register
	//case 2:
	default:
		return false
	}
	r := redisutils.NewRedisCline()
	c, err := r.Get(context.Background(), prefix+mobile).Result()
	if err != nil {
		return false
	}
	if c == code {
		// 删除redis数据
		r.Del(context.Background(), prefix+mobile)
		return true
	}

	return true
}
