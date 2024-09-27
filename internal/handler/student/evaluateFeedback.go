package student

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
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
		utils.JsonFail(c, 200506, "评价失败")
		return
	}
	utils.JsonSuccess(c, nil)
	
}
