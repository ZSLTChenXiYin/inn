package main

import (
	captcha "captcha/kitex_gen/captcha"
	"context"
	"captcha/biz/service"
)

// CaptchaServiceImpl implements the last service interface defined in the IDL.
type CaptchaServiceImpl struct{}

// EmailCaptcha implements the CaptchaServiceImpl interface.
func (s *CaptchaServiceImpl) EmailCaptcha(ctx context.Context, req *captcha.EmailCaptchaRequest) (resp *captcha.EmailCaptchaResponse, err error) {
	resp, err = service.NewEmailCaptchaService(ctx).Run(req)

	return resp, err
}

// EmailValidate implements the CaptchaServiceImpl interface.
func (s *CaptchaServiceImpl) EmailValidate(ctx context.Context, req *captcha.EmailValidateRequest) (resp *captcha.EmailValidateResponse, err error) {
	resp, err = service.NewEmailValidateService(ctx).Run(req)

	return resp, err
}

// EmailRefresh implements the CaptchaServiceImpl interface.
func (s *CaptchaServiceImpl) EmailRefresh(ctx context.Context, req *captcha.EmailRefreshRequest) (resp *captcha.EmailRefreshResponse, err error) {
	resp, err = service.NewEmailRefreshService(ctx).Run(req)

	return resp, err
}
