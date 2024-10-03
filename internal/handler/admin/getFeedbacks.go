package admin

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetFeedbacks(c *gin.Context) {
	feedbackList, err := service.GetFeedbacks_()
	if err != nil {
		zap.L().Error("获取反馈记录失败", zap.Error(err))
		return
	}
	if feedbackList == nil {
		utils.JsonFail(c, 200511, "当前无反馈记录")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"feedback_list": feedbackList,
	})
}
