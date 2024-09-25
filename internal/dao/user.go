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
