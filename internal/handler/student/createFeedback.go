package student

import (
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"

	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateFeedbackData struct {
	Title    string                  `form:"title" binding:"required"`
	Category int                     `form:"category" binding:"required"`
	IsUrgent bool                    `form:"is_urgent" binding:"required"`
	Name     string                  `form:"name"`
	Content  string                  `form:"content" binding:"required"`
	Images   []*multipart.FileHeader `form:"images" binding:"required"`
	Time     time.Time
}

func CreateFeedback(c *gin.Context) {
	var data CreateFeedbackData
	if err := c.ShouldBind(&data); err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	data.Time = time.Now()
	err := service.CreateFeedback(c.GetInt("user_id"), data.Title, data.Category, data.IsUrgent, data.Name, data.Content, data.Images, data.Time)
	if err != nil {
		utils.JsonFail(c, 200504, "发布失败")
		return
	}
	utils.JsonSuccess(c, nil)
}
