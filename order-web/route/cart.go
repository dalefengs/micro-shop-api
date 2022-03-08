package route

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/order-web/api/cart"
	"micro-shop-api/order-web/middleware"
)

func IniCartRoute(router *gin.RouterGroup) {
	orderRoute := router.Group("shop_cart").Use(middleware.JwtAuth())
	{
		orderRoute.GET("", cart.List)          //购物车列表
		orderRoute.DELETE("/:id", cart.Delete) //删除条目
		orderRoute.POST("", cart.New)          //添加商品到购物车
		orderRoute.PATCH("/:id", cart.Update)  //修改条目
	}
}
