package service

import (
	auth "auth/kitex_gen/auth"
	"context"
	"fmt"
	"time"

	"auth/biz/utils"
)

type TokenService struct {
	ctx context.Context
}

// NewTokenService 创建 TokenService 实例
func NewTokenService(ctx context.Context) *TokenService {
	return &TokenService{ctx: ctx}
}

// Run 为每个 ID 生成 Token
func (s *TokenService) Run(req *auth.TokenRequest) (resp *auth.TokenResponse, err error) {
	// 初始化一个切片，用于存储生成的 Token
	tokens := make([]*auth.Token, 0, len(req.Id))

	// 遍历 req.Id 数组
	for _, id := range req.Id {
		// 调用 GenerateToken 函数生成 Token
		access_token, err := utils.GenerateToken(id, time.Now().Add(time.Hour*4).Unix())
		if err != nil {
			// 如果生成 Token 失败，返回错误
			return nil, fmt.Errorf("failed to generate token for ID %s: %v", id, err)
		}

		refresh_token, err := utils.GenerateToken(id, time.Now().Add(time.Hour*24*7).Unix())
		if err != nil {
			return nil, fmt.Errorf("failed to generate token for ID %s: %v", id, err)
		}

		// 创建 Token 结构体实例，并将生成的 Token 赋值给 AccessToken 字段
		token := &auth.Token{
			AccessToken:  access_token, // 将生成的 Token 赋值给 AccessToken
			RefreshToken: refresh_token,
		}

		// 将生成的 Token 添加到 tokens 切片中
		tokens = append(tokens, token)
	}

	// 返回响应对象，包含生成的 Token 切片
	resp = &auth.TokenResponse{
		Tokens: tokens,
	}

	return
}
