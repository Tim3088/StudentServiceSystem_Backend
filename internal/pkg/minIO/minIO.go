package minIO

import (
	"StudentServiceSystem/internal/global"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

var MI *minio.Client

func Init() {
	// 初始化minio
	endpoint := global.Config.GetString("minio.endpoint")
	accessKeyID := global.Config.GetString("minio.accessKey")
	secretAccessKey := global.Config.GetString("minio.secretKey")
	useSSL := global.Config.GetBool("minio.useSSL")
	minioclient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		zap.L().Fatal("连接minIO数据库失败！", zap.Error(err))
	}
	MI = minioclient
}
