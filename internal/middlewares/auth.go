package middlewares

import (
	"gin--/internal/services"
	"gin--/internal/utils/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			logger.Logger.Warn("token 为空 未授权")
			return
		}
		// 验证 Token 格式是否正确
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			logger.Logger.Warn("token 格式错误 未授权")
			return
		}

		user := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(parts[1], user, services.Secret())

		if err != nil {
			logger.Logger.Warn("解析失败")
			return
		}
		if !token.Valid {
			logger.Logger.Warn("token 无效")
			return
		}

		c.Set("me", user)
		c.Next() //通过后继续处理请求

	}
}
