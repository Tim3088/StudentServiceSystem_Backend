package user

import (
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType int    `json:"user_type" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone"`
}

func Register(c *gin.Context) {
	// TODO

}
