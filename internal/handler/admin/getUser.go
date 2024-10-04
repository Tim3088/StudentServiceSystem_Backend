package admin

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetUserRequest struct {
	Username string `json:"username" binding:"required"`
}

func GetUser(c *gin.Context) {
	info, err := service.GetUserByUserID(c.GetInt("user_id"))
	if err != nil {
		zap.L().Error("获取管理员信息失败", zap.Error(err))
		return
	}
	if info.UserType != 3 {
		utils.JsonFail(c, 200513, "当前用户不是超级管理员")
		return
	}

	var data GetUserRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	user, err := service.GetUserByUsername(data.Username)
	if err != nil {
		utils.JsonFail(c, 200503, "用户不存在")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"name":     user.Name,
		"password": user.Password,
		"contact_info": gin.H{
			"email": user.Email,
			"phone": user.Phone,
		},
	})
}
