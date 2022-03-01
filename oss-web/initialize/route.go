package initialize

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/oss-web/middleware"
	"micro-shop-api/oss-web/route"
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
	ApiGroup := Router.Group("oss/v1")

	route.InitOssRoute(ApiGroup)

	return Router
}
