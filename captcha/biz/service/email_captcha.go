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
	// 初始化响应对象，Success 字段用于记录每个邮箱的验证码发送是否成功
	resp = &captcha.EmailCaptchaResponse{
		Success: make([]bool, len(req.Email)),
	}

	// 遍历请求中的每个邮箱
	for index, email := range req.Email {
		// 检查 Redis 中是否已经存在该邮箱的验证码
		result, _ := redis.RedisClient.Exists(s.ctx, email).Result()
		if result == 1 {
			// 如果邮箱已经存在，跳过当前邮箱的处理
			continue
		}

		// 生成随机验证码
		random_captcha := utils.GenerateCaptcha()

		// 将验证码存储到 Redis 中，键为邮箱地址，值为验证码
		err := redis.RedisClient.Set(s.ctx, email, random_captcha, 0).Err()
		if err != nil {
			// 如果存储失败，记录错误日志并跳过当前邮箱的处理
			klog.Error("redis hset captcha error: ", err)
			continue
		}

		// 设置验证码的过期时间为 5 分钟
		err = redis.RedisClient.Expire(s.ctx, email, 5*time.Minute).Err()
		if err != nil {
			// 如果设置过期时间失败，记录错误日志并跳过当前邮箱的处理
			klog.Error("redis hexpire captcha error: ", err)
			continue
		}

		// 发送验证码到邮箱
		err = utils.SendEmailCaptcha(email, random_captcha)
		if err != nil {
			// 如果发送失败，记录错误日志并跳过当前邮箱的处理
			klog.Error("send email captcha error: ", err)
			continue
		}

		// 标记当前邮箱的验证码发送成功
		resp.Success[index] = true
	}

	// 返回响应对象
	return
}
