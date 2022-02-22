package route

import (
	"github.com/gin-gonic/gin"
	"micro-shop-api/user-web/api"
	"micro-shop-api/user-web/middleware"
)

func InitUserRoute(router *gin.RouterGroup) {
	UserRoute := router.Group("user").Use(middleware.JwtAuth())
	{
		//UserRoute.GET("list", middleware.JwtAuth(), api.GetUserList)
		UserRoute.GET("list", api.GetUserList)
	}
}
