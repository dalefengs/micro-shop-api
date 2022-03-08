package route

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/order-web/api/order"
	"micro-shop-api/order-web/middleware"
)

func InitOrderRoute(router *gin.RouterGroup) {
	orderRoute := router.Group("order").Use(middleware.JwtAuth())
	{
		orderRoute.GET("", order.List)       // 订单列表
		orderRoute.POST("", order.New)       // 新建订单
		orderRoute.GET("/:id", order.Detail) // 订单详情
		orderRoute.GET("/notify")
	}
}
