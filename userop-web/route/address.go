package route

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/userop-web/api/address"
	"micro-shop-api/userop-web/middleware"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("address")
	{
		AddressRouter.GET("", middleware.JwtAuth(), address.List)          // 轮播图列表页
		AddressRouter.DELETE("/:id", middleware.JwtAuth(), address.Delete) // 删除轮播图
		AddressRouter.POST("", middleware.JwtAuth(), address.New)          //新建轮播图
		AddressRouter.PUT("/:id", middleware.JwtAuth(), address.Update)    //修改轮播图信息
	}
}
