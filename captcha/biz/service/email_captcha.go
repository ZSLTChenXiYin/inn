package service

import (
	"captcha/biz/dal/redis"
	"captcha/biz/utils"
	captcha "captcha/kitex_gen/captcha"
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

type EmailCaptchaService struct {
	ctx context.Context
} // NewEmailCaptchaService new EmailCaptchaService
func NewEmailCaptchaService(ctx context.Context) *EmailCaptchaService {
	return &EmailCaptchaService{ctx: ctx}
}

// Run create note info
func (s *EmailCaptchaService) Run(req *captcha.EmailCaptchaRequest) (resp *captcha.EmailCaptchaResponse, err error) {
	// Finish your business logic.

	resp = &captcha.EmailCaptchaResponse{
		Success: make([]bool, len(req.Email)),
	}

	for index, email := range req.Email {
		result, _ := redis.RedisClient.Exists(s.ctx, email).Result()
		if result == 1 {
			continue
		}

		random_captcha := utils.GenerateCaptcha()

		err := redis.RedisClient.Set(s.ctx, email, random_captcha, 0).Err()
		if err != nil {
			klog.Error("redis hset captcha error: ", err)
			continue
		}

		err = redis.RedisClient.Expire(s.ctx, email, 5*time.Minute).Err()
		if err != nil {
			klog.Error("redis hexpire captcha error: ", err)
			continue
		}

		err = utils.SendEmailCaptcha(email, random_captcha)
		if err != nil {
			klog.Error("send email captcha error: ", err)
			continue
		}

		resp.Success[index] = true
	}

	return
}
