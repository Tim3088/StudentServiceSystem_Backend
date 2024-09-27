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
	err := service.ReplyFeedback(data.FeedbackID, data.Reply)
	if err != nil {
		zap.L().Error("回复反馈失败", zap.Error(err))
		return
	}
	utils.JsonSuccess(c, nil)
}
