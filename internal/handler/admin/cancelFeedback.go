package admin

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CancelFeedbackRequest struct {
	FeedbackID int `json:"id" binding:"required"`
}

func CancelFeedback(c *gin.Context) {
	info, err := service.GetUserByUserID(c.GetInt("user_id"))
	if err != nil {
		zap.L().Error("获取管理员信息失败", zap.Error(err))
		return
	}

	if info.UserType != 2 && info.UserType != 3 {
		utils.JsonFail(c, 200512, "当前用户不是管理员")
		return
	}

	var data CancelFeedbackRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}

	feedback, err := service.FindFeedback(data.FeedbackID)
	if err != nil {
		utils.JsonFail(c, 200503, "反馈不存在")
		return
	}

	err = service.CancelFeedback(feedback.ID)
	if err != nil {
		zap.L().Error("撤销接单失败", zap.Error(err))
		return
	}

	utils.JsonSuccess(c, nil)
}
