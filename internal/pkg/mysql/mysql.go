package mysql

import (
	"StudentServiceSystem/internal/global"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	user := global.Config.GetString("mysql.user")
	pass := global.Config.GetString("mysql.pass")
	host := global.Config.GetString("mysql.host")
	port := global.Config.GetString("mysql.port")
	dbname := global.Config.GetString("mysql.dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		zap.L().Fatal("连接数据库失败！", zap.Error(err))
	}

	err = autoMigration(db)

	if err != nil {
		zap.L().Fatal("数据库迁移失败！", zap.Error(err))
	}
	return db
}
