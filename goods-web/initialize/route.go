package initialize

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/goods-web/middleware"
	"micro-shop-api/goods-web/route"
	"net/http"
)

// Routers 初始化路由
func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Cors()) // 跨域请求
	Router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"success": true,
		})
	})
	ApiGroup := Router.Group("/g/v1")

	route.InitGoodsRoute(ApiGroup)
	route.InitCommonRoute(ApiGroup)
	route.InitCategoryRouter(ApiGroup)
	route.InitBannerRouter(ApiGroup)

	return Router
}
