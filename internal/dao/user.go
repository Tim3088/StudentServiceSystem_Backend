package dao

import (
	"StudentServiceSystem/internal/model"
	"context"
)

func (d *Dao) GetUserByUsername(ctx context.Context, userName string) (*model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("username = ?", userName).First(&user).Error
	return &user, err
}

func (d *Dao) UpdateUser(ctx context.Context, username string, email string, phone string, password string) error {
	return d.orm.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).Updates(map[string]interface{}{
		"email":    email,
		"phone":    phone,
		"password": password,
	}).Error
}

func (d *Dao) CreateUser(ctx context.Context, user *model.User) error {
	return d.orm.WithContext(ctx).Create(user).Error
}
