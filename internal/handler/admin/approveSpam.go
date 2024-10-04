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
		utils.JsonFail(c, 200512, "获取管理员信息失败")
		return
	}
	if info.UserType != 3 {
		utils.JsonFail(c, 200512, "当前用户不是超级管理员")
		return
	}
	var request struct {
		FeedbackID string `json:"feedback_id" binding:"required"`
		Approval   bool   `json:"approval" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.JsonFail(c, 200513, "请求参数错误")
		return
	}
	if !request.Approval {
		utils.JsonSuccess(c, gin.H{
			"message": "审核已拒绝",
		})
		return
	}
	service.ApproveSpam(request.FeedbackID)
	utils.JsonSuccess(c, gin.H{
		"message": "反馈已成功审核标记为垃圾反馈",
	})
}
