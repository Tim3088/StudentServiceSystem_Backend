package service

import (
    "StudentServiceSystem/internal/model"
    "StudentServiceSystem/internal/pkg/minIO"
    "encoding/json"
    "mime/multipart"
    "time"
)

func CreateFeedback(userID int, title string, category int, isUrgent bool, name string, content string, images []*multipart.FileHeader, time time.Time) error {
    // 上传图片并获取 URL 列表
    imageURLs, err := minIO.UploadFile(images)
    if err != nil {
        return err
    }

    // 将 URL 列表序列化为 JSON 字符串
    imageURLsJSON, err := json.Marshal(imageURLs)
    if err != nil {
        return err
    }

    // 创建 Feedback 实例并存储到数据库
    return d.CreateFeedback(ctx,&model.Feedback{
        UserID:   userID,
        Title:    title,
        Time:     time,
        Category: category,
        IsUrgent: isUrgent,
        Name:     name,
        Content:  content,
        Images:   string(imageURLsJSON), // 存储 JSON 字符串
    })
}