package service

import "StudentServiceSystem/internal/model"

func GetUserByUsername(userName string) (*model.User, error) {
	return d.GetUserByUsername(ctx, userName)
}
