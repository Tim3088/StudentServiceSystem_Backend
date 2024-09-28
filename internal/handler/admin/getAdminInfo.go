package admin

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func GetAdminInfo(c *gin.Context) {
	info, err := service.GetAdminInfo(c.GetInt("user_id"))
	if err != nil {
		zap.L().Error("获取管理员信息失败", zap.Error(err))
		return
	}
	if info.UserType != 2 {
		utils.JsonFail(c, 200511, "当前用户不是管理员")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"username": info.Username,
		"name":     info.Name,
		"contact_info": gin.H{
			"email": info.Email,
			"phone": info.Phone,
		},
	})
}
