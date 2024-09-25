package user

import (
	"StudentServiceSystem/pkg/utils"
	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	utils.RefreshTokenHandler(c)
}
