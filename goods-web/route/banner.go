package route

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/goods-web/api/banners"
	"micro-shop-api/goods-web/middleware"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	//BannerRouter := Router.Group("banners").Use(middlewares.Trace())
	BannerRouter := Router.Group("banners")
	{
		BannerRouter.GET("", banners.List)                                // 轮播图列表页
		BannerRouter.DELETE("/:id", middleware.JwtAuth(), banners.Delete) // 删除轮播图
		BannerRouter.POST("", middleware.JwtAuth(), banners.New)          //新建轮播图
		BannerRouter.PUT("/:id", middleware.JwtAuth(), banners.Update)    //修改轮播图信息
	}
}
