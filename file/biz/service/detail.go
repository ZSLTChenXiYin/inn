package service

import (
	"context"
	"encoding/json"
	"file/biz/dal/model"
	"file/biz/dal/mysql"
	"file/biz/dal/query"
	"file/biz/dal/redis"
	file "file/kitex_gen/file"

	"github.com/cloudwego/kitex/pkg/klog"
)

type DetailService struct {
	ctx context.Context
} // NewDetailService new DetailService
func NewDetailService(ctx context.Context) *DetailService {
	return &DetailService{ctx: ctx}
}

// Run create note info
func (s *DetailService) Run(req *file.DetailRequest) (resp *file.DetailResponse, err error) {
	// Finish your business logic.

	// 创建响应对象
	resp = &file.DetailResponse{
		Files: make([]*file.DetailInfo, len(req.Files)),
	}

	// 获取数据库对象
	multfile_query := query.Use(mysql.DB)

	for index, detail := range req.Files {
		// 验证参数
		if detail.FileId == "" || detail.UserId == "" {
			continue
		}

		// 从Redis中获取公开文件信息
		json_data, _ := redis.RedisClient.HGet(s.ctx, redis.HASHMAP_OPEN_FILE, detail.FileId).Result()
		if json_data != "" {
			// 将从redis获取的json数据转换为结构体
			open_file := &model.File{}
			err := json.Unmarshal([]byte(json_data), open_file)
			if err == nil {
				resp.Files[index] = &file.DetailInfo{
					FileName:    open_file.FileName,
					OwnerId:     open_file.OwnerID,
					AccessLevel: uint32(open_file.AccessLevel),
					FileSize:    uint64(open_file.FileSize),
				}

				continue
			} else {
				// 文件信息反序列化失败，访问拒绝，进入查询数据库阶段
				klog.Error("unmarshal open file error: ", err)
			}
		}

		// 从数据库中查询文件信息
		query_file, err := multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.FileID.Eq(detail.FileId)).First()
		if err != nil {
			klog.Error("query file error: ", err)
			continue
		}

		resp.Files[index] = &file.DetailInfo{
			FileName:    query_file.FileName,
			OwnerId:     query_file.OwnerID,
			AccessLevel: uint32(query_file.AccessLevel),
			FileSize:    uint64(query_file.FileSize),
		}

		// 将文件信息存入redis
		if query_file.AccessLevel == 1 || query_file.AccessLevel == 2 {
			// 将文件信息序列化
			json_data, err := json.Marshal(query_file)
			if err != nil {
				klog.Error("marshal protected file error: ", err)
				continue
			}

			// 将文件信息存入redis
			err = redis.RedisClient.HSet(s.ctx, redis.HASHMAP_OPEN_FILE, query_file.FileID, string(json_data)).Err()
			if err != nil {
				klog.Error("redis hset open file error: ", err)
			}
		}
	}

	return
}
