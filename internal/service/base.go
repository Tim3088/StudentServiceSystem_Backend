package service

import (
	"StudentServiceSystem/internal/dao"
	"context"
	"gorm.io/gorm"
)

var (
	ctx = context.Background()
	d   *dao.Dao
)

func ServiceInit(db *gorm.DB) {
	d = dao.New(db)
}
