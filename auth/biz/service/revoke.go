package service

import (
	"auth/biz/dal/redis"
	"auth/biz/utils"
	"auth/kitex_gen/auth"
	"context"
	"fmt"
	"log"
	"time"
)

type RevokeService struct {
	ctx context.Context
}

// NewRevokeService 创建 RevokeService 实例
func NewRevokeService(ctx context.Context) *RevokeService {
	return &RevokeService{ctx: ctx}
}

// Run 处理 Token 撤销请求
func (s *RevokeService) Run(req *auth.RevokeRequest) (resp *auth.RevokeResponse, err error) {
	if len(req.AccessToken) == 0 {
		return nil, fmt.Errorf("no access tokens provided")
	}

	success := make([]bool, len(req.AccessToken))

	for i, token := range req.AccessToken {
		// 解析 Token，获取到期时间
		claims, err := utils.ParseToken(token)
		if err != nil {
			log.Printf("Failed to parse token: %v", err)
			success[i] = false
			continue
		}

		// 获取 JWT 过期时间
		endtime, ok := claims["time"].(float64)
		if !ok {
			log.Println("Token does not contain valid exp claim")
			success[i] = false
			continue
		}

		expireTime := time.Until(time.Unix(int64(endtime), 0)) // 计算 Token 剩余的存活时间

		// 将撤销的 Token 存入 Redis，设置与 JWT 过期时间一致
		err = redis.RedisClient.Set(s.ctx, token, "true", 0).Err()
		if err != nil {
			log.Printf("Failed to revoke token %s: %v", token, err)
			success[i] = false
			continue
		}

		err = redis.RedisClient.Expire(s.ctx, token, expireTime).Err()
		if err != nil {
			log.Printf("Failed to set expiration for token %s: %v", token, err)
			success[i] = false
			continue
		}

		success[i] = true
	}

	resp = &auth.RevokeResponse{
		Success: success,
	}
	return
}
