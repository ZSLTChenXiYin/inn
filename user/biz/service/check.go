package service

import (
	"context"
	"user/biz/dal/mysql"
	"user/biz/dal/query"
	user "user/kitex_gen/user"

	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckService struct {
	ctx context.Context
} // NewCheckService new CheckService
func NewCheckService(ctx context.Context) *CheckService {
	return &CheckService{ctx: ctx}
}

// Run create note info
func (s *CheckService) Run(req *user.CheckRequest) (resp *user.CheckResponse, err error) {
	// Finish your business logic.

	user_query := query.Use(mysql.DB)

	query_users, err := user_query.User.WithContext(s.ctx).Where(user_query.User.UserName.Eq(req.UserName)).Find()
	if err != nil {
		klog.Error("query user error: ", err)
		return
	}

	query_users_len := len(query_users)
	if query_users_len == 0 {
		return
	}

	resp = &user.CheckResponse{
		UserId: make([]string, query_users_len),
	}

	for index, query_user := range query_users {
		resp.UserId[index] = query_user.UserID
	}

	return
}
