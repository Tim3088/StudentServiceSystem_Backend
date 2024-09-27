package student

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetStudentInfo(c *gin.Context) {
	studentInfo, err := service.GetStudentInfo(c.GetInt("user_id"))
	if err != nil {
		zap.L().Error("获取学生信息失败", zap.Error(err))
		return
	}
	utils.JsonSuccess(c, gin.H{
		"username": studentInfo.Username,
		"name":     studentInfo.Name,
		"contact_info": gin.H{
			"email": studentInfo.Email,
			"phone": studentInfo.Phone,
		},
	})
}
