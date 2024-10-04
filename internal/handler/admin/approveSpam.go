package admin

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ApproveSpam 审核标记为垃圾反馈
func ApproveSpam(c *gin.Context) {
	info, err := service.GetUserByUserID(c.GetInt("user_id"))
	if err != nil {
		zap.L().Error("获取管理员信息失败", zap.Error(err))
		return
	}
	if info.UserType != 3 {
		utils.JsonFail(c, 200513, "当前用户不是超级管理员")
		return
	}
	var request struct {
		FeedbackID int `json:"id" binding:"required"`
		Approval   int `json:"approval" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	if request.Approval == 2 {
		utils.JsonSuccess(c, nil)
		return
	}
	service.ApproveSpam(request.FeedbackID)
	utils.JsonSuccess(c, nil)
}
