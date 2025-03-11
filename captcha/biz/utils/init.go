package utils

import (
	"captcha/conf"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func Init() {
	email_captcha_pool, err := email.NewPool(
		fmt.Sprintf("%s:%d", conf.GetConf().Email.SmtpHost, conf.GetConf().Email.SmtpPort),
		10,
		smtp.PlainAuth(
			"",
			conf.GetConf().Email.SmtpUsername,
			conf.GetConf().Email.SmtpPassword,
			conf.GetConf().Email.SmtpHost,
		),
	)
	if err != nil {
		panic(err)
	}
	emailCaptchaPool = email_captcha_pool
}
