package redis

import (
	"StudentServiceSystem/internal/global"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func Init() {
	pass := global.Config.GetString("redis.pass")
	host := global.Config.GetString("redis.host")
	port := global.Config.GetString("redis.port")
	db := global.Config.GetInt("redis.db")
	// 初始化连接
	Rdb = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: pass,
		DB:       db,
	})
}
