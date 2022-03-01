package route

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/oss-web/api"
)

func InitOssRoute(router *gin.RouterGroup) {
	ossRoute := router.Group("oss")
	{
		ossRoute.GET("token", api.Token)
		ossRoute.POST("callback", api.Callback)
	}
}
