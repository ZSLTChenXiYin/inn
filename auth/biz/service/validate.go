package service

import (
	"auth/biz/dal/redis"
	"auth/biz/utils"
	"auth/kitex_gen/auth"
	"context"
	"fmt"
)

type ValidateService struct {
	ctx context.Context
}

// NewValidateService 创建 ValidateService 实例
func NewValidateService(ctx context.Context) *ValidateService {
	return &ValidateService{ctx: ctx}
}

// Run 处理 Token 验证请求
func (s *ValidateService) Run(req *auth.ValidateRequest) (resp *auth.ValidateResponse, err error) {
	if len(req.AccessToken) == 0 {
		return nil, fmt.Errorf("no access tokens provided")
	}

	ids := make([]string, 0, len(req.AccessToken))

	// 检查 Token 是否已被撤销
	for _, token := range req.AccessToken {
		revoked, err := redis.RedisClient.Get(s.ctx, token).Result()
		if err == nil && revoked == "true" {
			continue // 如果 Token 已被撤销，跳过
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			continue // 如果 Token 无效，跳过
		}

		id, ok := claims["id"].(string)
		if !ok {
			continue // Token 不包含有效 ID
		}

		ids = append(ids, id)
	}

	resp = &auth.ValidateResponse{
		Id: ids,
	}
	return
}
