package service

import (
	"context"
	"encoding/json"
	"file/biz/dal/model"
	"file/biz/dal/mysql"
	"file/biz/dal/query"
	"file/biz/dal/redis"
	file "file/kitex_gen/file"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

type DeleteService struct {
	ctx context.Context
} // NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *file.DeleteRequest) (resp *file.DeleteResponse, err error) {
	// Finish your business logic.

	// 创建响应对象
	resp = &file.DeleteResponse{
		Success: make([]bool, len(req.Files)),
	}

	// 获取数据库对象
	multfile_query := query.Use(mysql.DB)

	for index, delete := range req.Files {
		// 验证参数
		if delete.FileId == "" || delete.UserId == "" {
			continue
		}

		// 从Redis中获取公开文件信息
		json_data, _ := redis.RedisClient.HGet(s.ctx, redis.HASHMAP_OPEN_FILE, delete.FileId).Result()
		// 如果Redis中不存在公开文件信息，则从数据库中获取
		if json_data == "" {
			// 从数据库中获取文件信息
			query_file, err := multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.FileID.Eq(delete.FileId)).First()
			if err != nil {
				klog.Error("query file error: ", err)
				continue
			}

			// 判断用户是否为文件所有者
			if delete.UserId != query_file.OwnerID {
				continue
			}
		} else {
			// 将从redis获取的json数据转换为结构体
			open_file := &model.File{}
			err := json.Unmarshal([]byte(json_data), open_file)
			if err != nil {
				klog.Error("json unmarshal error: ", err)

				// 从数据库中获取文件信息
				open_file, err = multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.FileID.Eq(delete.FileId)).First()
				if err != nil {
					klog.Error("query file error: ", err)
					continue
				}
			}

			// 判断用户是否为文件所有者
			if delete.UserId != open_file.OwnerID {
				continue
			}
		}

		// 从数据库删除文件
		_, err := multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.FileID.Eq(delete.FileId)).Delete()
		if err != nil {
			klog.Error("delete file error: ", err)
			resp.Success[index] = false
		}

		resp.Success[index] = true

		// 删除redis中的文件信息
		err = redis.RedisClient.HDel(s.ctx, redis.HASHMAP_OPEN_FILE, delete.FileId).Err()
		if err != nil {
			klog.Error("redis hdel open file error: ", err)
		}

		// 延时删除redis中的文件信息
		go func() {
			time.Sleep(time.Second * 1)
			err = redis.RedisClient.HDel(s.ctx, redis.HASHMAP_OPEN_FILE, delete.FileId).Err()
			if err != nil {
				klog.Error("redis hdel open file error: ", err)
			}
		}()
	}

	return
}
