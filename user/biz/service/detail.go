package service

import (
	"context"
	"user/biz/dal/mysql"
	"user/biz/dal/query"
	user "user/kitex_gen/user"

	"github.com/cloudwego/kitex/pkg/klog"
)

type DetailService struct {
	ctx context.Context
} // NewDetailService new DetailService
func NewDetailService(ctx context.Context) *DetailService {
	return &DetailService{ctx: ctx}
}

// Run create note info
func (s *DetailService) Run(req *user.DetailRequest) (resp *user.DetailResponse, err error) {
	// Finish your business logic.

	// 创建响应对象
	resp = &user.DetailResponse{
		Details: make([]*user.DetailInfo, len(req.UserId)),
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

		resp.Details[index] = &user.DetailInfo{
			UserName:  query_user.UserName,
			Email:     query_user.Email,
			Signature: query_user.PersonalizedSignature,
			Picture:   query_user.Picture,
			CreatedAt: query_user.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return
}
