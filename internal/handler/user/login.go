package user

import (
	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	// TODO
}
