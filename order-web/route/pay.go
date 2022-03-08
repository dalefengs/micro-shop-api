package route

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/order-web/api/pay"
)

func InitPayOrderRoute(router *gin.RouterGroup) {
	orderRoute := router.Group("pay")
	{
		orderRoute.POST("/alipay/notify", pay.Notify)
	}
}
