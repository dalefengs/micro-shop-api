package status

// 0 - 2000 公用状态码
// 2000 - 3000 用户服务
// 3000 - 4000 商品
//4000 - 5000 订单

var (
	// 系统状态码
	OK                        = Code{0, "SUCCESS"}
	ServerError               = Code{1000, "服务器异常"}
	Fail                      = Code{1001, "操作失败"}
	InvalidOperation          = Code{1002, "无效操作"}
	InvalidParameter          = Code{1003, "参数错误"}
	UnavailableServer         = Code{1010, "服务不可用"}
	AuthFail                  = Code{1011, "身份验证失败"}
	AuthExpired               = Code{1012, "您还未登录"}
	SmsSendCodeFail           = Code{1013, "短信发送失败"}
	CodeIncorrect             = Code{1014, "验证码不正确"}
	RegisterCenterConnectFail = Code{1015, "连接注册中心失败"}
	ServerNotFind             = Code{1016, "服务丢失或未找到"}
	NotFound                  = Code{1017, "不存在"}
	ResourceExhausted         = Code{1018, "资源不足"}
	InnerError                = Code{1019, "内部异常"}

	NotFoundUser  = Code{1020, "用户不存在"}
	PasswordError = Code{1021, "账号或密码错误"}
	AlreadyExists = Code{1022, "用户已存在"}
)

type Code struct {
	Code int
	Msg  string
}
