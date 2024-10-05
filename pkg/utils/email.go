package utils

import (
	"StudentServiceSystem/internal/global"
	"context"
	"strconv"

	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

var UGoemail = uGoemail{}

type uGoemail struct{}

func (e *uGoemail) SendMail(ctx context.Context, to, cc, subject, content string) (result int, err error) {
	//获取配置文件信息
	host := global.Config.GetString("goemail.host")
	port, _ := strconv.Atoi(global.Config.GetString("goemail.port"))
	username := global.Config.GetString("goemail.username")
	password := global.Config.GetString("goemail.password")

	//发送邮件
	m := gomail.NewMessage()

	m.SetHeader("From", username) //发送邮箱

	m.SetHeader("To", to) //主送

	m.SetHeader("Subject", subject) //标题

	m.SetBody("text/html", content) // 发送html格式邮件，发送的内容

	d := gomail.NewDialer(host, port, username, password)

	if err = d.DialAndSend(m); err != nil {
		zap.L().Error("邮件发送失败", zap.Error(err))
		return 2, err
	}
	return 1, nil
}
