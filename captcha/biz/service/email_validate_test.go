package service

import (
	"context"
	"testing"
	captcha "captcha/kitex_gen/captcha"
)

func TestEmailValidate_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEmailValidateService(ctx)
	// init req and assert value

	req := &captcha.EmailValidateRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
