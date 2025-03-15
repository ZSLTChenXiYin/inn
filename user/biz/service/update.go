package service

import (
	"context"
	"user/biz/dal/mysql"
	"user/biz/dal/query"
	"user/biz/utils"
	user "user/kitex_gen/user"

	"github.com/cloudwego/kitex/pkg/klog"
)

type UpdateService struct {
	ctx context.Context
} // NewUpdateService new UpdateService
func NewUpdateService(ctx context.Context) *UpdateService {
	return &UpdateService{ctx: ctx}
}

// Run create note info
func (s *UpdateService) Run(req *user.UpdateRequest) (resp *user.UpdateResponse, err error) {
	// Finish your business logic.

	// 创建响应对象
	resp = &user.UpdateResponse{
		Success: make([]bool, len(req.Updates)),
	}

	// 获取数据库对象
	user_query := query.Use(mysql.DB)

	for index, update := range req.Updates {
		// 验证参数
		if update.UserId == "" {
			continue
		}

		// 从数据库中获取用户信息
		query_user, err := user_query.User.WithContext(s.ctx).Where(user_query.User.UserID.Eq(update.UserId)).First()
		if err != nil {
			klog.Error("get user error: ", err)
			continue
		}

		seccess := true

		if update.NewUserName != "" {
			// 更新用户名
			_, err := user_query.User.WithContext(s.ctx).Where(user_query.User.UserID.Eq(update.UserId)).Update(user_query.User.UserName, update.NewUserName)
			if err != nil {
				klog.Error("update user name error: ", err)
				seccess = false
			}
		}

		if update.NewPassword != "" {
			// 生成密码哈希
			password_hash, err := utils.GenerateHashPassword(update.NewPassword)
			if err != nil {
				klog.Error("bcrypt error: ", err)
				seccess = false
			}

			// 更新密码哈希
			_, err = user_query.User.WithContext(s.ctx).Where(user_query.User.UserID.Eq(update.UserId)).Update(user_query.User.UserPasswordHash, password_hash)
			if err != nil {
				klog.Error("update password error: ", err)
				seccess = false
			}
		}

		if update.NewPicture != "" {
			// 更新头像
			_, err := user_query.User.WithContext(s.ctx).Where(user_query.User.UserID.Eq(update.UserId)).Update(user_query.User.Picture, update.NewPicture)
			if err != nil {
				klog.Error("update picture error: ", err)
				seccess = false
			}
		}

		if update.NewSignature != "" {
			// 更新个性签名
			_, err := user_query.User.WithContext(s.ctx).Where(user_query.User.UserID.Eq(update.UserId)).Update(user_query.User.PersonalizedSignature, update.NewSignature)
			if err != nil {
				klog.Error("update personalized signature error: ", err)
				seccess = false
			}
		}

		// 如果更新失败，则回滚
		if !seccess {
			_, err := user_query.User.WithContext(s.ctx).Where(user_query.User.UserID.Eq(update.UserId)).Update(user_query.User.ALL, query_user)
			if err != nil {
				klog.Error("rollback error: ", err)
			}
		}

		resp.Success[index] = true
	}

	return
}
