package test

import (
	"captcha/kitex_gen/captcha"
	"captcha/kitex_gen/captcha/captchaservice"
	"context"
	"testing"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func TestCaptcha(t *testing.T) {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		panic(err)
	}

	cli, err := captchaservice.NewClient("captcha", client.WithResolver(r))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	email_captcha_resp, err := cli.EmailCaptcha(ctx, &captcha.EmailCaptchaRequest{
		Email: []string{"1234567@qq.com"},
	})
	if err != nil {
		panic(err)
	}

	t.Log(email_captcha_resp)

	email_refresh_resp, err := cli.EmailRefresh(ctx, &captcha.EmailRefreshRequest{
		Email: []string{"1234567@qq.com"},
	})
	if err != nil {
		panic(err)
	}

	t.Log(email_refresh_resp)

	/*
		email_validate_resp, err := cli.EmailValidate(ctx, &captcha.EmailValidateRequest{
			Lists: []*captcha.EmailValidateList{
				{
					Email:   "1234567@qq.com",
					Captcha: "123456",
				},
			},
		})

		t.Log(email_validate_resp)
	*/
}
