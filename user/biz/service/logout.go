package service

import (
	"context"
	"user/biz/dal/mysql"
	"user/biz/dal/query"
	"user/biz/dal/redis"
	user "user/kitex_gen/user"

	"github.com/cloudwego/kitex/pkg/klog"
)

type LogoutService struct {
	ctx context.Context
} // NewLogoutService new LogoutService
func NewLogoutService(ctx context.Context) *LogoutService {
	return &LogoutService{ctx: ctx}
}

// Run create note info
func (s *LogoutService) Run(req *user.LogoutRequest) (resp *user.LogoutResponse, err error) {
	// Finish your business logic.

	// 创建响应对象
	resp = &user.LogoutResponse{
		Success: make([]bool, len(req.UserId)),
	}

	// 获取数据库对象
	user_query := query.Use(mysql.DB)

	for index, user_id := range req.UserId {
		// 验证参数
		if user_id == "" {
			continue
		}

		// 从数据库中获取用户信息
		query_user, err := user_query.User.WithContext(s.ctx).Where(user_query.User.UserID.Eq(user_id)).First()
		if err != nil {
			klog.Error("get user error: ", err)
			continue
		}

		// 删除用户信息
		_, err = user_query.User.WithContext(s.ctx).Where(user_query.User.UserID.Eq(user_id)).Delete()
		if err != nil {
			klog.Error("delete user error: ", err)
			continue
		}

		// 删除redis中用户信息
		err = redis.RedisClient.Del(s.ctx, query_user.Email).Err()
		if err != nil {
			klog.Error("redis hdel user error: ", err)
		}

		resp.Success[index] = true
	}

	return
}
