package dao

import (
	"StudentServiceSystem/internal/model"
	"StudentServiceSystem/internal/pkg/minIO"
	"context"

	"go.uber.org/zap"
)

func (d *Dao) CreateFeedback(ctx context.Context, feedback *model.Feedback) error {
	return d.orm.WithContext(ctx).Create(feedback).Error
}

func (d *Dao) GetFeedbacks(ctx context.Context, userID int) ([]map[string]interface{}, error) {
	var feedbackList []model.Feedback
	err := d.orm.WithContext(ctx).Where("user_id = ?", userID).Find(&feedbackList).Error
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
		var status int = 0
		//如果ReceiverID不为0，说明已被接收
		if feedback.ReceiverID != 0 {
			status = 1
		}
		res = append(res, map[string]interface{}{
			"id":         feedback.ID,
			"title":      feedback.Title,
			"time":       feedback.Time,
			"category":   feedback.Category,
			"is_urgent":  feedback.IsUrgent,
			"content":    feedback.Content,
			"images":     imageUrls,
			"reply":      feedback.Reply,
			"evaluation": feedback.Evaluation,
			"status":     status,
		})
	}
	return res, nil
}

func (d *Dao) EvaluateFeedback(ctx context.Context, feedbackID int, evaluation string) error {
	return d.orm.WithContext(ctx).Model(&model.Feedback{}).Where("id = ?", feedbackID).Update("evaluation", evaluation).Error
}

func (d *Dao) GetStudentInfo(ctx context.Context, userID int) (model.User, error) {
	var user model.User
	err := d.orm.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
