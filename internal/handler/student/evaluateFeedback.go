package student

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EvaluateFeedbackData struct {
	FeedbackID int    `json:"id" binding:"required"`
	Evaluation string `json:"evaluation" binding:"required"`
}

func EvaluateFeedback(c *gin.Context) {
	var data EvaluateFeedbackData
	if err := c.ShouldBind(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	err := service.EvaluateFeedback(data.FeedbackID, data.Evaluation)
	if err != nil {
		zap.L().Error("评价反馈失败", zap.Error(err))
		return
	}
	utils.JsonSuccess(c, nil)
	
}
