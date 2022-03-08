package initialize

import (
	"fmt"
	"github.com/go-playground/locales/en"
	en2 "github.com/go-playground/validator/v10/translations/en"
	"micro-shop-api/userop-web/global"
	validator2 "micro-shop-api/userop-web/validator"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
)

// InitTrans 初始化验证器
func InitTrans(locale string) (err error) {
	// 修改 gin框架中的Validator引擎，实现定制
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}
	//注册一个获取Json tag的自定义方法
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	zhT := zh.New() //中文翻译器
	enT := en.New() //英文翻译器
	//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
	uni := ut.New(enT, zhT, enT)
	global.Trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s)", locale)
	}
	switch locale {
	case "en":
		_ = en2.RegisterDefaultTranslations(v, global.Trans)
	case "zh":
		_ = zh2.RegisterDefaultTranslations(v, global.Trans)
	default:
		_ = zh2.RegisterDefaultTranslations(v, global.Trans)
	}

	// 自定义验证器
	err = v.RegisterValidation("mobile", validator2.ValidateMobile)
	if err != nil {
		return err
	}
	err = registerValidator(v)
	return err
}

// 注册自定义验证器 - 待优化
func registerValidator(v *validator.Validate) error {
	err := v.RegisterValidation("mobile", validator2.ValidateMobile)
	if err != nil {
		return err
	}
	err = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
		return ut.Add("mobile", "手机号格式不正确！", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("mobile", fe.Field())
		return t
	})
	return err
}
