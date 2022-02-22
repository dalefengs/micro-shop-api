package route

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/user-web/api"
)

func InitCommonRoute(Route *gin.RouterGroup) {
	commonRoute := Route.Group("c")
	{
		commonRoute.POST("send_code", api.SendSmsCode)
		commonRoute.POST("pwd_login", api.PasswordLogin)
		commonRoute.POST("register", api.RegisterUser)
		commonRoute.GET("captcha", api.Captcha)
	}
}
