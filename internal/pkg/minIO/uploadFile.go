package minIO

import (
	"StudentServiceSystem/internal/global"
	"context"
	"mime/multipart"
	"net/url"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)
func getFile(objectName string) (string, error) {
    ctx := context.Background()
    bucketName := global.Config.GetString("minio.bucketName")

    // 设置过期时间为1小时
    reqParams := url.Values{}
    presignedURL, err := MI.PresignedGetObject(ctx, bucketName, objectName, time.Hour, reqParams)
    if err != nil {
        zap.L().Error("Failed to generate presigned URL.", zap.Error(err))
        return "", err
    }

    return presignedURL.String(), nil
}
func UploadFile(images []*multipart.FileHeader) ([]string, error) {
	ctx := context.Background()
	bucketName := global.Config.GetString("minio.bucketName")

	var uploadedFiles []string

	for _, fileHeader := range images {
		// 打开文件
		file, err := fileHeader.Open()
		if err != nil {
			zap.L().Error("Failed to open file.", zap.Error(err))
			return nil, err
		}
		defer file.Close()

		// 上传文件
		objectName := filepath.Base(fileHeader.Filename)
		_, err = MI.PutObject(ctx, bucketName, objectName, file, fileHeader.Size, minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")})
		if err != nil {
			zap.L().Error("Failed to upload file.", zap.Error(err))
			return nil, err
		}

		// 构建文件 URL
		fileURL,err := getFile(objectName)
		if err != nil {
			zap.L().Error("Failed to get file URL.", zap.Error(err))
			return nil, err
		}
		uploadedFiles = append(uploadedFiles, fileURL)
	}

	return uploadedFiles, nil
}
