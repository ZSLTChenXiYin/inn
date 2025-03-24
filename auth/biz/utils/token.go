package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const JWTSecret = "KyJNv$&b)e(%QFSn7kn6"

// GenerateToken 根据 ID 生成 JWT Token
func GenerateToken(id string, time int64) (string, error) {
	// 创建一个新的 JWT Token，使用 HS256 签名方法
	token := jwt.New(jwt.SigningMethodHS256)

	// 获取 Token 的 Claims 部分
	claims := token.Claims.(jwt.MapClaims)

	// 设置 Token 的 Claims
	claims["id"] = id     // 用户 ID
	claims["time"] = time // 过期时间

	// 使用常量密钥对 Token 进行签名，生成最终的 Token 字符串
	jwtSecret := []byte(JWTSecret) // JWT 密钥
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}
