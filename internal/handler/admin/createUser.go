package admin

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"
	"regexp"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	UserType int    `json:"user_type" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone"`
}
func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
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
	re := regexp.MustCompile(`^\d+$`)
	if !re.MatchString(req.Username) {
		utils.JsonFail(c, 200502, "用户名必须为纯数字")
		return
	}
	if len(req.Password) < 8 || len(req.Password) > 16 {
		utils.JsonFail(c, 200503, "密码长度必须在8-16位")
		return
	}
	if req.UserType != 1 && req.UserType != 2 {
		utils.JsonFail(c, 200504, "用户类型错误")
		return
	}
	_, err = service.GetUserByUsername(req.Username)
	if err == nil {
		utils.JsonFail(c, 200505, "用户名已存在")
		return
	}
	if req.Phone != "" && len(req.Phone) != 11 {
		utils.JsonFail(c, 200506, "手机号必须为11位")
		return
	}
	err = service.Register(req.Username, req.Name, req.Password, req.UserType, req.Email, req.Phone)
	if err != nil {
		zap.L().Error("创建用户失败", zap.Error(err))
		return
	}
	utils.JsonSuccess(c, nil)
}
