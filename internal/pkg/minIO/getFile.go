package minIO

import (
	"StudentServiceSystem/internal/global"
	"context"
	"net/url"
	"time"

	"go.uber.org/zap"
)

// getFile 生成预签名 URL
func GetFile(objectName string) (string, error) {
    ctx := context.Background()
    bucketName := global.Config.GetString("minio.bucketName")

    // 设置过期时间为1小时
    reqParams := url.Values{}
    reqParams.Set("response-content-disposition", "attachment; filename=\""+objectName+"\"")
    presignedURL, err := MI.PresignedGetObject(ctx, bucketName, objectName, time.Hour, reqParams)
    if err != nil {
        zap.L().Error("Failed to generate presigned URL.", zap.Error(err))
        return "", err
    }

    return presignedURL.String(), nil
}