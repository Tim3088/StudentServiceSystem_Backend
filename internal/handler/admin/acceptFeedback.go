package admin

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AcceptFeedbackRequest struct {
	FeedbackID int `json:"id" binding:"required"`
}
func AcceptFeedback(c *gin.Context) {
	var data AcceptFeedbackRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	err:=service.FindFeedback(data.FeedbackID)
	if err != nil {
		utils.JsonFail(c, 200503, "反馈不存在")
		return
	}
	err = service.AcceptFeedback(data.FeedbackID,c.GetInt("user_id"))
	utils.JsonSuccess(c, nil)
}
