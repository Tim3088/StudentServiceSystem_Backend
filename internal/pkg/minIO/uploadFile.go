package minIO

import (
	"StudentServiceSystem/internal/global"
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

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

		// 将文件名添加到上传文件列表
		uploadedFiles = append(uploadedFiles, objectName)
	}

	return uploadedFiles, nil
}
