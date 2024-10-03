package dao

import (
	"StudentServiceSystem/internal/model"
	"context"

)

func (d *Dao) ReplyFeedback(ctx context.Context,feedbackID int, reply string) error {
	return d.orm.Model(&model.Feedback{}).Where("id = ?", feedbackID).Update("reply", reply).Error
}

func (d *Dao) GetAdminInfo(ctx context.Context,userID int) (model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (d *Dao) Update(ctx context.Context, username string, name string, userType int, newUsername string, password string) {
    // 使用 Updates 方法一次性更新多个字段
    d.orm.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).Updates(map[string]interface{}{
        "username":  newUsername,
        "name":      name,
        "user_type": userType,
        "password":  password,
    })
}
func (d *Dao) MarkFeedback(ctx context.Context,feedbackID int) error {
	return d.orm.Model(&model.ReportFeedback{}).Create(&model.ReportFeedback{FeedbackID: feedbackID}).Error
}

func (d *Dao) FindReportFeedback(ctx context.Context,feedbackID int) error {
	var reportFeedback model.ReportFeedback
	err := d.orm.WithContext(ctx).Where("feedback_id = ?", feedbackID).First(&reportFeedback).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) FindFeedback(ctx context.Context,feedbackID int) error {
	var feedback model.Feedback
	err := d.orm.WithContext(ctx).Where("id = ?", feedbackID).First(&feedback).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) AcceptFeedback(ctx context.Context,feedbackID int,userID int) error {
	return d.orm.Model(&model.Feedback{}).Where("id = ?", feedbackID).Update("receiver_id", userID).Error
}