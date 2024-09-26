package middleware

import (
	"StudentServiceSystem/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func HandleNotFound(c *gin.Context) {
	// 记录未找到的请求信息
	zap.L().Warn("未找到请求路径", zap.String("path", c.Request.URL.Path))

	// 返回404 Not Found响应
	utils.JsonResponse(c, 404, 200404, http.StatusText(http.StatusNotFound), nil)
}
