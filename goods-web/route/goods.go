package route

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/goods-web/api/goods"
	"micro-shop-api/goods-web/middleware"
)

func InitGoodsRoute(router *gin.RouterGroup) {
	goodsRoute := router.Group("goods").Use(middleware.JwtAuth())
	{
		goodsRoute.GET("", goods.List)
		goodsRoute.POST("", goods.New, middleware.JwtAuth())
		goodsRoute.GET("/:id", goods.Detail)
		goodsRoute.DELETE("/:id", goods.Delete)
		goodsRoute.PATCH("/:id", goods.UpdateStatus)
		goodsRoute.PUT("/:id", goods.Update)
	}
}
