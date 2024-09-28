package admin

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"
	"regexp"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateUserData struct {
	Username    string `json:"username" binding:"required"`
	Name        string `json:"name"`
	UserType    int    `json:"user_type"`
	NewUsername string `json:"new_username"`
	Password    string `json:"password"`
}

func UpdateUser(c *gin.Context) {
	var data UpdateUserData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	info,err:=service.GetUserByUserID(c.GetInt("user_id"))
	if err != nil {
		zap.L().Error("获取管理员信息失败", zap.Error(err))
		return
	}
	if info.UserType != 3 {
		utils.JsonFail(c, 200511, "当前用户不是超级管理员")
		return
	}
	_, err = service.GetUserByUsername(data.Username)
	if err != nil {
		utils.JsonFail(c, 200506, "用户不存在")
		return
	}
	re := regexp.MustCompile(`^\d+$`)
	if !re.MatchString(data.NewUsername) {
		utils.JsonFail(c, 200502, "用户名必须为纯数字")
		return
	}
	if len(data.Password) < 8 || len(data.Password) > 16 {
		utils.JsonFail(c, 200503, "密码长度必须在8-16位")
		return
	}
	if data.UserType != 1 && data.UserType != 2 {
		utils.JsonFail(c, 200504, "用户类型错误")
		return
	}
	service.Update(data.Username,data.Name, data.UserType, data.NewUsername, data.Password)
	utils.JsonSuccess(c, nil)
}
