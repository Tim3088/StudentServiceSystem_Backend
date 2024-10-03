package admin

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	feedback,err:=service.FindFeedback(data.FeedbackID)
	if err != nil {
		utils.JsonFail(c, 200503, "反馈不存在")
		return
	}
	err = service.AcceptFeedback(data.FeedbackID,c.GetInt("user_id"))
	if err != nil {
		zap.L().Error("接单失败", zap.Error(err))
		return
	}
	info,err := service.GetStudentInfo(feedback.UserID)
	if err != nil {
		zap.L().Error("获取学生信息失败", zap.Error(err))
		return
	}
	
	utils.JsonSuccess(c, gin.H{
		"name": info.Name,
		"contact_info": gin.H{
			"email": info.Email,
			"phone": info.Phone,
		},
	})

}
