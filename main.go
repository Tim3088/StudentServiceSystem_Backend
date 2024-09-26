package main

import (
	"StudentServiceSystem/internal/global"
	"StudentServiceSystem/internal/middleware"
	"StudentServiceSystem/internal/pkg/minIO"
	"StudentServiceSystem/internal/pkg/mysql"
	"StudentServiceSystem/internal/pkg/redis"
	"StudentServiceSystem/internal/router"
	"StudentServiceSystem/internal/service"
	"StudentServiceSystem/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	//初始化日志
	debug := global.Config.GetBool("debug")
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	utils.Init(debug)

	//初始化数据库
	db := mysql.Init()
	service.ServiceInit(db)
	redis.Init()
	minIO.Init()

	//初始化路由
	r := gin.Default()
	r.NoMethod(middleware.HandleNotFound) // 中间件统一处理404
	r.NoRoute(middleware.HandleNotFound)
	router.Init(r)

	//启动服务器
	err := r.Run()
	if err != nil {
		zap.L().Fatal("Server start failed", zap.Error(err))
	} else {
		zap.L().Info("Server start success")
	}
}
