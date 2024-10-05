package admin

import (
	"StudentServiceSystem/pkg/utils"
	"StudentServiceSystem/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ReplyFeedbackData struct {
	FeedbackID int    `json:"id" binding:"required"`
	Reply      string `json:"reply" binding:"required"`
}
func ReplyFeedback(c *gin.Context) {
	var data ReplyFeedbackData
	if err := c.ShouldBind(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	feedback, err := service.FindFeedback(data.FeedbackID)
	if err != nil {
		utils.JsonFail(c, 200503, "反馈不存在")
		return
	}
	userInfo, err := service.GetStudentInfo(feedback.UserID)
	if err != nil {
		zap.L().Error("获取学生信息失败", zap.Error(err))
		return
	}
	err = service.ReplyFeedback(data.FeedbackID, data.Reply,userInfo.Email)
	if err != nil {
		zap.L().Error("回复反馈失败", zap.Error(err))
		return
	}
	utils.JsonSuccess(c, nil)
}
