package router

import (
	"StudentServiceSystem/internal/handler/admin"
	"StudentServiceSystem/internal/handler/student"
	"StudentServiceSystem/internal/handler/user"
	"StudentServiceSystem/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const api = "/api"
	apiGroup := r.Group(api)
	{
		userGroup := apiGroup.Group("/user")
		{
			userGroup.GET("/refresh", user.RefreshToken)
			userGroup.POST("/reg", user.Register)
			userGroup.POST("/login", user.Login)
			userGroup.PUT("/:username", user.UpdateUser).Use(middleware.JwtAuthMiddleware())
		}
		studentGroup := apiGroup.Group("/student")
		{
			studentGroup.GET("", student.GetStudentInfo)
			feedbackGroup := studentGroup.Group("/feedback").Use(middleware.JwtAuthMiddleware())
			{
				feedbackGroup.GET("", student.GetFeedback)
				feedbackGroup.POST("", student.CreateFeedback)
				feedbackGroup.PUT("", student.EvaluateFeedback)
			}
		}
		adminGroup := apiGroup.Group("/admin")
		{
			adminGroup.GET("", admin.GetAdminInfo)
			feedbackGroup := adminGroup.Group("/feedback").Use(middleware.JwtAuthMiddleware())
			{
				feedbackGroup.GET("", admin.GetFeedbacks)
				feedbackGroup.GET("/approve-spam", admin.GetSpamFeedbacks)
				feedbackGroup.POST("", admin.MarkFeedback)
				feedbackGroup.POST("/approve-spam", admin.ApproveSpam)
				feedbackGroup.PUT("/accept", admin.AcceptFeedback)
				feedbackGroup.PUT("/cancel", admin.CancelFeedback)
				feedbackGroup.PUT("/reply", admin.ReplyFeedback)
			}
			usersGroup := adminGroup.Group("/users").Use(middleware.JwtAuthMiddleware())
			{
				usersGroup.GET("", admin.GetUser)
				usersGroup.POST("", admin.CreateUser)
				usersGroup.PUT("", admin.UpdateUser)
				usersGroup.DELETE("", admin.DeleteUser)
			}
		}

	}
}
