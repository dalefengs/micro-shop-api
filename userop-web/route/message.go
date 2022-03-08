package route

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/userop-web/api/message"
	"micro-shop-api/userop-web/middleware"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message").Use(middleware.JwtAuth())
	{
		MessageRouter.GET("", message.List) // 轮播图列表页
		MessageRouter.POST("", message.New) //新建轮播图
	}
}
