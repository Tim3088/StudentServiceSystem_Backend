package user

import (
	"StudentServiceSystem/pkg/utils"
	"StudentServiceSystem/internal/service"
	"regexp"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType int    `json:"user_type" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone"`
}

func Register(c *gin.Context) {
	var data RegisterRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	re := regexp.MustCompile(`^\d+$`)
	if !re.MatchString(data.Username) {
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
	_, err := service.GetUserByUsername(data.Username)
	if err == nil {
		utils.JsonFail(c, 200505, "用户名已存在")
		return
	}
	if data.Phone != "" && len(data.Phone) != 11 {
		utils.JsonFail(c, 200506, "手机号必须为11位")
		return
	}
	err = service.Register(data.Username, data.Name, data.Password, data.UserType, data.Email, data.Phone)
	if err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		return
	}
	utils.JsonSuccess(c, nil)
}
