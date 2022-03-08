package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"micro-shop-api/order-web/extends/jwtauth"
	"micro-shop-api/order-web/global/response"
	status2 "micro-shop-api/order-web/global/status"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 header 中的 token
		token := c.GetHeader("Authorization")
		j := jwtauth.NewJWT()
		// 验证 token
		claims, err := j.ParseToken(token)
		if err != nil {
			zap.S().Infof("登录错误：%s", err.Error())
			response.FailResponse(c, http.StatusUnauthorized, status2.AuthExpired, "", "")
			c.Abort()
		}
		c.Set("user_id", claims.Id)
		c.Set("claims", claims)
		c.Next()
	}
}
