package service

import "StudentServiceSystem/internal/model"

func ReplyFeedback(feedbackID int, reply string) error {
	return d.ReplyFeedback(ctx, feedbackID, reply)
}

func GetAdminInfo(userID int) (model.User, error) {
	return d.GetAdminInfo(ctx, userID)
}

func Update(username string, name string, userType int, newUsername string, password string) {
	d.Update(ctx, username, name, userType, newUsername, password)
}

func MarkFeedback(feedbackID int) error {
	return d.MarkFeedback(ctx, feedbackID)
}

func FindReportFeedback(feedbackID int) error {
	return d.FindReportFeedback(ctx, feedbackID)
}

func FindFeedback(feedbackID int) error {
	return d.FindFeedback(ctx, feedbackID)
}

func AcceptFeedback(feedbackID int, userID int) error {
	return d.AcceptFeedback(ctx, feedbackID, userID)
}
