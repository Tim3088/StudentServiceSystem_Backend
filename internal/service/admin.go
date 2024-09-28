package service

import "StudentServiceSystem/internal/model"

func ReplyFeedback(feedbackID int, reply string) error {
	return d.ReplyFeedback(ctx, feedbackID, reply)
}

func GetAdminInfo(userID int) (model.User, error) {
	return d.GetAdminInfo(ctx, userID)
}