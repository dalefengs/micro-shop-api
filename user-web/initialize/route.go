package initialize

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/user-web/middleware"
	"micro-shop-api/user-web/route"
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
	ApiGroup := Router.Group("/u/v1")

	route.InitUserRoute(ApiGroup)
	route.InitCommonRoute(ApiGroup)

	return Router
}
