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