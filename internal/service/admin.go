package service

import (
	"StudentServiceSystem/internal/model"
	"StudentServiceSystem/pkg/utils"

	"go.uber.org/zap"
)

func ReplyFeedback(feedbackID int, reply string,email string) error {
	mail, _ := utils.UGoemail.SendMail(ctx, email, "", "反馈处理结果", "您的反馈已经处理完毕，请及时查看")
	zap.L().Info("邮件发送结果", zap.Int("result", mail))
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

func FindFeedback(feedbackID int) (model.Feedback, error) {
	return d.FindFeedback(ctx, feedbackID)
}

func AcceptFeedback(feedbackID int, userID int) error {
	return d.AcceptFeedback(ctx, feedbackID, userID)
}

func GetFeedbacks_() ([]map[string]interface{}, error) {
	return d.GetFeedbacks_(ctx)
}

func DeleteUser(userID int) {
	d.DeleteUser(ctx, userID)
}

func GetSpamFeedbacks() ([]map[string]interface{}, error) {
	return d.GetSpamFeedbacks(ctx)
}

func ApproveSpam(feedbackID int, email string) {
	mail, _ := utils.UGoemail.SendMail(ctx, email, "", "反馈处理结果", "您的反馈已经处理完毕，请及时查看")
	zap.L().Info("邮件发送结果", zap.Int("result", mail))
	d.ApproveSpam(ctx, feedbackID)
}

func CancelFeedback(feedbackID int) error {
	return d.CancelFeedback(ctx, feedbackID)
}
