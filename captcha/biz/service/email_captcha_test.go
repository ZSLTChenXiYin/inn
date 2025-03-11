package service

import (
	"context"
	"testing"
	captcha "captcha/kitex_gen/captcha"
)

func TestEmailCaptcha_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEmailCaptchaService(ctx)
	// init req and assert value

	req := &captcha.EmailCaptchaRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
