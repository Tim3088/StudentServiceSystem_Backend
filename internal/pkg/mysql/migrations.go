package mysql

import (
	"StudentServiceSystem/internal/model"
	"gorm.io/gorm"
)

func autoMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
		model.Feedback{},
		model.ReportFeedback{},
	)
}
