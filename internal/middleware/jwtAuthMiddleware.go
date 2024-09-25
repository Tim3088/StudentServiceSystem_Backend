package middleware

import (
	"StudentServiceSystem/internal/pkg/redis"
	"StudentServiceSystem/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"strings"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		const bearerPrefix = "Bearer "
		// 提取实际的令牌
		tokenString := strings.TrimPrefix(authHeader, bearerPrefix)
		mc, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.JsonFail(c, 200510, "无效的访问令牌")
			c.Abort()
			return
		}
		if mc.TokenType != "access" {
			utils.JsonFail(c, 200510, "无效的访问令牌")
			c.Abort()
			return
		}
		// 检查Redis中是否存在该令牌
		ctx := context.Background()
		_, err = redis.Rdb.Get(ctx, "access_token:"+tokenString).Result()
		if err != nil {
			utils.JsonFail(c, 200510, "无效的访问令牌")
			c.Abort()
			return
		}
		c.Set("user_id", mc.UserID)
		c.Next()
	}
}
