package service

import "StudentServiceSystem/internal/model"

func GetUserByUsername(userName string) (*model.User, error) {
	return d.GetUserByUsername(ctx, userName)
}
func GetUserByUserID(userID int) (*model.User, error) {
	return d.GetUserByUserID(ctx, userID)
}
func UpdateUser(username string, email string, phone string, password string) {
	d.UpdateUser(ctx, username, email, phone, password)
}

func Register(username string, name string, password string, usertype int, email string, phone string) error {
	return d.CreateUser(ctx, &model.User{
		Username: username,
		Name:     name,
		Password: password,
		UserType: usertype,
		Email:    email,
		Phone:    phone,
	})
}
