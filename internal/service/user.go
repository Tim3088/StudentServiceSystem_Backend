package service

import "StudentServiceSystem/internal/model"

func GetUserByUsername(userName string) (*model.User, error) {
	return d.GetUserByUsername(ctx, userName)
}
func UpdateUser(username string, email string, phone string, password string) {
	d.UpdateUser(ctx, username, email, phone, password)
}