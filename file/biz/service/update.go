package service

import (
	"context"
	"file/biz/dal/mysql"
	"file/biz/dal/query"
	"file/biz/dal/redis"
	"file/biz/utils"
	file "file/kitex_gen/file"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

type UpdateService struct {
	ctx context.Context
} // NewUpdateService new UpdateService
func NewUpdateService(ctx context.Context) *UpdateService {
	return &UpdateService{ctx: ctx}
}

// Run create note info
func (s *UpdateService) Run(req *file.UpdateRequest) (resp *file.UpdateResponse, err error) {
	// Finish your business logic.

	// 创建响应对象
	resp = &file.UpdateResponse{
		Success: make([]bool, len(req.Files)),
	}

	// 获取数据库对象
	multfile_query := query.Use(mysql.DB)

	for index, update := range req.Files {
		// 验证参数
		if update.FileId == "" || update.UserId == "" {
			continue
		}

		// 从数据库中获取文件信息
		query_file, err := multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.FileID.Eq(update.FileId)).First()
		if err != nil {
			klog.Error("get file info error: ", err)
			continue
		}

		// 验证用户
		if update.UserId != query_file.OwnerID {
			continue
		}

		success := true

		// 更新文件信息
		if update.NewAccessLevel != 0 {
			// 更新文件访问权限
			_, err := multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.FileID.Eq(update.FileId)).Update(multfile_query.File.AccessLevel, update.NewAccessLevel)
			if err != nil {
				klog.Error("update access level error: ", err)
				success = false
			}

			// 更新文件密码
			if update.NewAccessLevel == 2 {
				if update.NewFilePassword != "" {
					password_hash, err := utils.GenerateHashPassword(update.NewFilePassword)
					if err != nil {
						klog.Error("bcrypt error: ", err)
						success = false
					}

					_, err = multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.FileID.Eq(update.FileId)).Update(multfile_query.File.FilePasswordHash, password_hash)
					if err != nil {
						klog.Error("update file password error: ", err)
						success = false
					}
				}
			}
		}

		if update.NewFileName != "" {
			// 更新文件名
			_, err := multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.FileID.Eq(update.FileId)).Update(multfile_query.File.FileName, update.NewFileName)
			if err != nil {
				klog.Error("update file name error: ", err)
				success = false
			}
		}

		if update.NewFilePath != "" {
			// 更新文件路径
			_, err := multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.FileID.Eq(update.FileId)).Update(multfile_query.File.FilePath, update.NewFilePath)
			if err != nil {
				klog.Error("update file path error: ", err)
				success = false
			}
		}

		if !success {
			// 回滚
			_, err := multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.FileID.Eq(update.FileId)).Update(multfile_query.File.ALL, query_file)
			if err != nil {
				klog.Error("rollback error: ", err)
				continue
			}
		}

		resp.Success[index] = true

		// 删除redis中的文件信息
		err = redis.RedisClient.HDel(s.ctx, redis.HASHMAP_OPEN_FILE, update.FileId).Err()
		if err != nil {
			klog.Error("redis hdel open file error: ", err)
		}

		// 延时删除redis中的文件信息
		go func() {
			time.Sleep(time.Second * 1)
			err = redis.RedisClient.HDel(s.ctx, redis.HASHMAP_OPEN_FILE, update.FileId).Err()
			if err != nil {
				klog.Error("redis hdel open file error: ", err)
			}
		}()
	}

	return
}
