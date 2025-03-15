package service

import (
	"context"
	"user/biz/dal/model"
	"user/biz/dal/mysql"
	"user/biz/dal/query"
	"user/biz/dal/redis"
	"user/biz/utils"
	user "user/kitex_gen/user"

	"github.com/cloudwego/kitex/pkg/klog"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// Finish your business logic.

	// 创建响应对象
	resp = &user.LoginResponse{
		UserId: make([]string, len(req.Logins)),
	}

	// 获取数据库对象
	user_query := query.Use(mysql.DB)

	for index, login := range req.Logins {
		// 验证参数
		if login.Email == "" || login.Password == "" {
			continue
		}

		// 从redis中获取用户id
		result_uuid, err := redis.RedisClient.Get(s.ctx, login.Email).Result()
		if err != nil {
			klog.Error("redis get user_id error: ", err)
		}

		var query_user *model.User
		if result_uuid == "" {
			// 如果redis中没有用户id，使用用户邮箱查询用户信息
			query_user, err = user_query.User.WithContext(s.ctx).Where(user_query.User.Email.Eq(login.Email)).First()
			if err != nil {
				klog.Error("get user error: ", err)
				continue
			}

			// 将用户id存入redis中
			err = redis.RedisClient.Set(s.ctx, login.Email, query_user.UserID, 0).Err()
			if err != nil {
				klog.Error("redis hset user error: ", err)
			}
		} else {
			// 如果redis中有用户id，使用用户id查询用户信息
			query_user, err = user_query.User.WithContext(s.ctx).Where(user_query.User.UserID.Eq(result_uuid)).First()
			if err != nil {
				klog.Error("get user error: ", err)
				continue
			}
		}

		// 验证密码
		err = utils.ValidatePassword(login.Password, query_user.UserPasswordHash)
		if err != nil {
			klog.Error("validate password error: ", err)
			continue
		}

		resp.UserId[index] = result_uuid
	}

	return
}
