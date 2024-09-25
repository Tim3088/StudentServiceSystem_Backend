package user

import (
	"StudentServiceSystem/pkg/utils"
	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var data LoginData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	user, err := service.GetUserByUsername(data.Username)
	if err != nil {
		utils.JsonFail(c, 200506, "用户不存在")
		return
	}
	if user.Password != data.Password {
		utils.JsonFail(c, 200507, "密码错误")
		return
	}
	//登陆成功 通过jwt生成token签名并响应
	tokenString, refreshTokenString, err := utils.GenerateTokens(user.ID)
	if err != nil {
		utils.JsonFail(c, 200508, "生成token失败")
		return
	}
	utils.JsonSuccess(c, gin.H{
		"access_token":  tokenString,
		"refresh_token": refreshTokenString,
		"user_type":     user.UserType,
		"contact_info": gin.H{
			"email": user.ContactInfo.Email,
			"phone": user.ContactInfo.Phone,
		},
	})
}
