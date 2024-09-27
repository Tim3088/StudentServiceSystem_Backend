package dao

import "StudentServiceSystem/internal/model"

func (d *Dao) ReplyFeedback(feedbackID int, reply string) error {
	return d.orm.Model(&model.Feedback{}).Where("id = ?", feedbackID).Update("reply", reply).Error
}