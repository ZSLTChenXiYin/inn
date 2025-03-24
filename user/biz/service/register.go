package service

import (
	"context"
	"time"
	"user/biz/dal/model"
	"user/biz/dal/mysql"
	"user/biz/dal/query"
	"user/biz/dal/redis"
	"user/biz/utils"
	user "user/kitex_gen/user"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// Finish your business logic.

	// 创建响应对象
	resp = &user.RegisterResponse{
		UserId: make([]string, len(req.Registers)),
	}

	// 获取数据库对象
	user_query := query.Use(mysql.DB)

	for index, register := range req.Registers {
		// 验证参数
		if register.UserName == "" || register.Email == "" || register.Password == "" {
			continue
		}

		// 从数据库中获取用户信息
		_, err := user_query.User.WithContext(s.ctx).Where(user_query.User.Email.Eq(register.Email)).First()
		if err == nil {
			klog.Error("email already exists")
			continue
		}

		// 生成uuid
		uuidv4 := uuid.New().String()

		// 生成密码哈希
		password_hash, err := utils.GenerateHashPassword(register.Password)
		if err != nil {
			klog.Error("bcrypt error: ", err)
			continue
		}

		// 创建用户
		new_user := &model.User{
			UserID:           uuidv4,
			UserName:         register.UserName,
			Email:            register.Email,
			UserPasswordHash: password_hash,
		}

		// 将用户信息插入数据库
		err = user_query.User.WithContext(s.ctx).Create(new_user)
		if err != nil {
			klog.Error("create user error: ", err)
			continue
		}

		// 将用户信息插入redis
		err = redis.RedisClient.Set(s.ctx, register.Email, new_user.UserID, 0).Err()
		if err != nil {
			klog.Error("redis hset user error: ", err)
		}

		// 设置用户信息过期时间
		err = redis.RedisClient.Expire(s.ctx, register.Email, 48*time.Hour).Err()
		if err != nil {
			klog.Error("redis hexpire user error: ", err)
		}

		resp.UserId[index] = new_user.UserID
	}

	return
}
