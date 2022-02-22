package forms

// SendCodeForm 短信验证码
type SendCodeForm struct {
	MobileForm
	CodeType int `json:"type" mapstructure:"binding:required"` // 1注册 2找回密码
}
