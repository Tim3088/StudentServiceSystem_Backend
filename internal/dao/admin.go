package dao

import (
	"StudentServiceSystem/internal/model"
	"StudentServiceSystem/internal/pkg/minIO"
	"context"

	"go.uber.org/zap"
)

func (d *Dao) ReplyFeedback(ctx context.Context, feedbackID int, reply string) error {
	return d.orm.Model(&model.Feedback{}).Where("id = ?", feedbackID).Update("reply", reply).Error
}

func (d *Dao) GetAdminInfo(ctx context.Context, userID int) (model.User, error) {
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

func (d *Dao) MarkFeedback(ctx context.Context, feedbackID int) error {
	return d.orm.Model(&model.ReportFeedback{}).Create(&model.ReportFeedback{FeedbackID: feedbackID}).Error
}

func (d *Dao) FindReportFeedback(ctx context.Context, feedbackID int) error {
	var reportFeedback model.ReportFeedback
	err := d.orm.WithContext(ctx).Where("feedback_id = ?", feedbackID).First(&reportFeedback).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) FindFeedback(ctx context.Context, feedbackID int) (model.Feedback, error) {
	var feedback model.Feedback
	err := d.orm.WithContext(ctx).Where("id = ?", feedbackID).First(&feedback).Error
	if err != nil {
		return model.Feedback{}, err
	}
	return feedback, nil
}

func (d *Dao) AcceptFeedback(ctx context.Context, feedbackID int, userID int) error {
	return d.orm.Model(&model.Feedback{}).Where("id = ?", feedbackID).Update("receiver_id", userID).Error
}

func (d *Dao) GetFeedbacks_(ctx context.Context) ([]map[string]interface{}, error) {
	var feedbackList []model.Feedback
	err := d.orm.WithContext(ctx).Find(&feedbackList).Error
	if err != nil {
		return nil, err
	}
	var res []map[string]interface{}
	for _, feedback := range feedbackList {
		// 反序列化 Images 字段
		images, err := feedback.GetImages()
		if err != nil {
			return nil, err
		}

		var imageUrls []string
		for _, imageName := range images {
			url, err := minIO.GetFile(imageName)
			if err != nil {
				zap.L().Error("Failed to get file.", zap.Error(err))
				return nil, err
			}
			imageUrls = append(imageUrls, url)
		}
		res = append(res, map[string]interface{}{
			"id":          feedback.ID,
			"title":       feedback.Title,
			"time":        feedback.Time,
			"category":    feedback.Category,
			"is_urgent":   feedback.IsUrgent,
			"content":     feedback.Content,
			"images":      imageUrls,
			"reply":       feedback.Reply,
			"evaluation":  feedback.Evaluation,
			"receiver_id": feedback.ReceiverID,
		})
	}
	return res, nil
}

func (d *Dao) DeleteUser(ctx context.Context, userID int) error {
	result := d.orm.WithContext(ctx).Delete(&model.User{}, userID)
	return result.Error
}

func (d *Dao) GetSpamFeedbacks(ctx context.Context) ([]map[string]interface{}, error) {
	var feedbackList []int
	err := d.orm.WithContext(ctx).Model(&model.ReportFeedback{}).Pluck("feedback_id", &feedbackList).Error
	if err != nil {
		return nil, err
	}
	var res []map[string]interface{}
	for _, feedbackID := range feedbackList {
		feedback, err := d.FindFeedback(ctx, feedbackID)
		if err != nil {
			return nil, err
		}
		images, err := feedback.GetImages()
		if err != nil {
			return nil, err
		}
		var imageUrls []string
		for _, imageName := range images {
			url, err := minIO.GetFile(imageName)
			if err != nil {
				zap.L().Error("Failed to get file.", zap.Error(err))
				return nil, err
			}
			imageUrls = append(imageUrls, url)
		}
		res = append(res, map[string]interface{}{
			"id":      feedback.ID,
			"title":   feedback.Title,
			"content": feedback.Content,
			"images":  imageUrls,
		})
	}
	return res, nil
}

func (d *Dao) UpdateReceiverID(ctx context.Context, feedbackID int, receiverID int) {
	d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("id = ?", feedbackID).Updates(map[string]interface{}{
		"receiver_id": receiverID,
	})
}

func (d *Dao) ApproveSpam(ctx context.Context, feedbackID int) {
	// 第一步
	d.UpdateReceiverID(ctx, feedbackID, -1)
	// 第二步
	message := "请您在提交反馈时确保内容的有效性和准确性，感谢您的理解和配合。如有异议，请重新反馈。"
	d.ReplyFeedback(ctx, feedbackID, message)
}

func (d *Dao) CancelFeedback(ctx context.Context, feedbackID int) error {
	d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("id = ?", feedbackID).Updates(map[string]interface{}{
		"receiver_id": 0,
	})
	return nil
}
