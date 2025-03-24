package service

import (
	"auth/biz/utils"
	auth "auth/kitex_gen/auth"
	"context"
	"fmt"
	"time"
)

type RefreshService struct {
	ctx context.Context
}

// NewRefreshService 创建 RefreshService 实例
func NewRefreshService(ctx context.Context) *RefreshService {
	return &RefreshService{ctx: ctx}
}

// Run 处理刷新 Token 的请求
func (s *RefreshService) Run(req *auth.RefreshRequest) (resp *auth.RefreshResponse, err error) {
	if len(req.RefreshToken) == 0 {
		return nil, fmt.Errorf("no refresh tokens provided")
	}

	accessTokens := make([]string, len(req.RefreshToken))

	for i, refreshToken := range req.RefreshToken {
		// 解析并验证 Refresh Token
		claims, err := utils.ParseToken(refreshToken)
		if err != nil {
			return nil, fmt.Errorf("invalid refresh token: %v", err)
		}

		id, ok := claims["id"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid token claims: missing 'id'")
		}

		// 生成新的 Access Token
		newAccessToken, err := utils.GenerateToken(id, time.Now().Add(time.Hour*4).Unix())
		if err != nil {
			return nil, fmt.Errorf("failed to generate access token: %v", err)
		}

		accessTokens[i] = newAccessToken
	}

	resp = &auth.RefreshResponse{
		AccessToken: accessTokens,
	}
	return
}
