package model

import (
	"encoding/json"
	"time"
)

type Feedback struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Title      string    `json:"title"`
	Time       time.Time `json:"time"`
	Category   int       `json:"category"`
	IsUrgent   bool      `json:"is_urgent"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	Images     string    `json:"images"` // 修改为 string 类型
	Reply      string    `json:"reply"`
	Evaluation string    `json:"evaluation"`
	ReceiverID int       `json:"receiver_id"`
}

// 将 Images 字段序列化为 JSON 字符串
func (f *Feedback) SetImages(images []string) error {
	data, err := json.Marshal(images)
	if err != nil {
		return err
	}
	f.Images = string(data)
	return nil
}

// 将 JSON 字符串反序列化为 []string
func (f *Feedback) GetImages() ([]string, error) {
	var images []string
	err := json.Unmarshal([]byte(f.Images), &images)
	return images, err
}
