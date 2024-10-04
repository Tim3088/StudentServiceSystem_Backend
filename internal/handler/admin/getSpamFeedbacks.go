package admin

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetSpamFeedbacks(c *gin.Context) {
	info, err := service.GetUserByUserID(c.GetInt("user_id"))
	if err != nil {
		zap.L().Error("获取管理员信息失败", zap.Error(err))
		return
	}
	if info.UserType != 3 {
		utils.JsonFail(c, 200512, "当前用户不是超级管理员")
		return
	}
	reportList, err := service.GetSpamFeedbacks()
	if err != nil {
		zap.L().Error("获取反馈记录失败", zap.Error(err))
		return
	}
	if reportList == nil {
		utils.JsonFail(c, 200511, "当前无被标记的反馈记录")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"report_list": reportList,
	})
}
