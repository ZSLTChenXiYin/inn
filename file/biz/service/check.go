package service

import (
	"context"
	"file/biz/dal/mysql"
	"file/biz/dal/query"
	file "file/kitex_gen/file"

	"github.com/cloudwego/kitex/pkg/klog"
)

type CheckService struct {
	ctx context.Context
} // NewCheckService new CheckService
func NewCheckService(ctx context.Context) *CheckService {
	return &CheckService{ctx: ctx}
}

// Run create note info
func (s *CheckService) Run(req *file.CheckRequest) (resp *file.CheckResponse, err error) {
	// Finish your business logic.

	resp = &file.CheckResponse{}

	multfile_query := query.Use(mysql.DB)

	query_files, err := multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.OwnerID.Eq(req.OwnerId)).Find()
	if err != nil {
		klog.Error("query file error: ", err)
		return
	}

	if req.UserId == req.OwnerId {
		resp.FileId = make([]string, len(query_files))

		for index, query_file := range query_files {
			resp.FileId[index] = query_file.FileID
		}
	} else {
		for _, query_file := range query_files {
			if query_file.AccessLevel == 1 || query_file.AccessLevel == 2 {
				resp.FileId = append(resp.FileId, query_file.FileID)
			}
		}
	}

	return
}
