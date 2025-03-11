package utils

import (
	"captcha/conf"
	"fmt"
	"time"

	"github.com/jordan-wright/email"
)

var (
	emailCaptchaPool *email.Pool
)

func SendEmailCaptcha(target_email string, captcha string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", "客栈", conf.GetConf().Email.SmtpUsername)
	e.To = []string{target_email}
	e.Subject = "客栈验证码"
	e.Text = []byte(fmt.Sprintf("您的验证码为：%s", captcha))

	return emailCaptchaPool.Send(e, 10*time.Second)
}
