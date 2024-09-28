package user

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UpdateData struct {
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func UpdateUser(c *gin.Context) {
	username := c.Param("username")
	var data UpdateData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	
	if len(data.Password) < 8 || len(data.Password) > 16 {
		utils.JsonFail(c, 200503, "密码长度必须在8-16位")
		return
	}
	
	if data.Phone != "" && len(data.Phone) != 11 {
		utils.JsonFail(c, 200506, "手机号必须为11位")
		return
	}
	
	service.UpdateUser(username,data.Email,data.Phone,data.Password)
	utils.JsonSuccess(c, nil)
}
