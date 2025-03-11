package service

import (
	"captcha/biz/dal/redis"
	"captcha/biz/utils"
	captcha "captcha/kitex_gen/captcha"
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

type EmailRefreshService struct {
	ctx context.Context
} // NewEmailRefreshService new EmailRefreshService
func NewEmailRefreshService(ctx context.Context) *EmailRefreshService {
	return &EmailRefreshService{ctx: ctx}
}

// Run create note info
func (s *EmailRefreshService) Run(req *captcha.EmailRefreshRequest) (resp *captcha.EmailRefreshResponse, err error) {
	// Finish your business logic.

	resp = &captcha.EmailRefreshResponse{
		Success: make([]bool, len(req.Email)),
	}

	for index, email := range req.Email {
		random_captcha := utils.GenerateCaptcha()

		err = redis.RedisClient.Set(s.ctx, email, random_captcha, 0).Err()
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
