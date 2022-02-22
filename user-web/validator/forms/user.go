package forms

type MobileForm struct {
	Mobile string `json:"mobile" binding:"required,mobile"`
}
type PasswordForm struct {
	Password string `json:"password" binding:"required,min=3,max=20"`
}

// PasswordLoginForm 密码登录
type PasswordLoginForm struct {
	MobileForm
	PasswordForm
}

// RegisterUserForm 用户注册
type RegisterUserForm struct {
	MobileForm
	PasswordForm
	SmsCode  string `json:"code" binding:"required"`
	Nickname string `json:"nickname" binding:"required,min=3,max=20"`
}
