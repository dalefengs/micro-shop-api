package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidateMobile 验证手机号
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[579]|15[0-3,5-9]|16[6]|17[0135678]|18[0-9]|19[89])\d{8}$`, mobile)
	return ok
}
