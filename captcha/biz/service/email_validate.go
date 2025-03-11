package service

import (
	"captcha/biz/dal/redis"
	captcha "captcha/kitex_gen/captcha"
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
)

type EmailValidateService struct {
	ctx context.Context
} // NewEmailValidateService new EmailValidateService
func NewEmailValidateService(ctx context.Context) *EmailValidateService {
	return &EmailValidateService{ctx: ctx}
}

// Run create note info
func (s *EmailValidateService) Run(req *captcha.EmailValidateRequest) (resp *captcha.EmailValidateResponse, err error) {
	// Finish your business logic.

	resp = &captcha.EmailValidateResponse{
		Success: make([]bool, len(req.Lists)),
	}

	for index, email_validate_list := range req.Lists {
		captcha, err := redis.RedisClient.Get(s.ctx, email_validate_list.Email).Result()
		if err != nil {
			klog.Error("redis hget captcha error: ", err)
			continue
		}

		if captcha == email_validate_list.Captcha {
			resp.Success[index] = true

			err = redis.RedisClient.Del(s.ctx, email_validate_list.Email).Err()
			if err != nil {
				klog.Error("redis hdel captcha error: ", err)
			}
		}
	}

	return
}
