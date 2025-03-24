package service

import (
	"context"
	"encoding/json"
	"file/biz/dal/model"
	"file/biz/dal/mysql"
	"file/biz/dal/query"
	"file/biz/dal/redis"
	"file/biz/utils"
	file "file/kitex_gen/file"

	"github.com/cloudwego/kitex/pkg/klog"
)

type DownloadService struct {
	ctx context.Context
} // NewDownloadService new DownloadService
func NewDownloadService(ctx context.Context) *DownloadService {
	return &DownloadService{ctx: ctx}
}

// Run create note info
func (s *DownloadService) Run(req *file.DownloadRequest) (resp *file.DownloadResponse, err error) {
	// Finish your business logic.

	// 创建响应对象
	resp = &file.DownloadResponse{
		Files: make([]*file.DownloadInfo, len(req.Lists)),
	}

	// 获取数据库对象
	multfile_query := query.Use(mysql.DB)

	for index, download_list := range req.Lists {
		// 验证参数
		if download_list.FileId == "" || download_list.UserId == "" {
			continue
		}

		// 从redis中获取文件信息
		json_data, _ := redis.RedisClient.HGet(s.ctx, redis.HASHMAP_OPEN_FILE, download_list.FileId).Result()
		if json_data != "" {
			// 将文件信息反序列化
			open_file := &model.File{}
			err := json.Unmarshal([]byte(json_data), open_file)
			if err == nil {
				// 判断用户是否为文件拥有者，如果是，则允许访问
				if download_list.UserId != open_file.OwnerID {
					switch open_file.AccessLevel {
					case 0:
						// 文件访问权限为私有，访问拒绝
						continue
					case 2:
						// 文件访问权限为加密，需要验证密码
						if open_file.FilePasswordHash != "" {
							err = utils.ValidatePassword(download_list.FilePassword, open_file.FilePasswordHash)
							if err != nil {
								klog.Error("validate password error: ", err)
							}
						}

						fallthrough
					case 1:
						// 文件访问权限为公开，允许访问

						resp.Files[index] = &file.DownloadInfo{
							FileName: open_file.FileName,
							FilePath: open_file.FilePath,
							FileSize: uint64(open_file.FileSize),
							CheckSum: open_file.CheckSum,
						}

						continue
					default:
						// 文件访问权限为未知，访问拒绝，进入查询数据库阶段
						klog.Error("access level error")
					}
				}
			} else {
				// 文件信息反序列化失败，访问拒绝，进入查询数据库阶段
				klog.Error("unmarshal open file error: ", err)
			}
		}

		// 从数据库中查询文件信息
		query_file, err := multfile_query.File.WithContext(s.ctx).Where(multfile_query.File.FileID.Eq(download_list.FileId)).First()
		if err != nil {
			klog.Error("get file info error: ", err)
			continue
		}

		// 判断用户是否为文件拥有者，如果是，则允许访问
		if download_list.UserId != query_file.OwnerID {
			switch query_file.AccessLevel {
			case 0:
				// 文件访问权限为私有，访问拒绝
				continue
			case 2:
				// 文件访问权限为加密，需要验证密码
				if query_file.FilePasswordHash != "" {
					err = utils.ValidatePassword(download_list.FilePassword, query_file.FilePasswordHash)
					if err != nil {
						klog.Error("validate password error: ", err)
						continue
					}
				}

				fallthrough
			case 1:
				// 文件访问权限为公开，允许访问

				// 将文件信息序列化
				json_data, err := json.Marshal(query_file)
				if err != nil {
					klog.Error("marshal protected file error: ", err)
					break
				}

				// 将文件信息存入redis
				err = redis.RedisClient.HSet(s.ctx, redis.HASHMAP_OPEN_FILE, download_list.FileId, string(json_data)).Err()
				if err != nil {
					klog.Error("redis hset open file error: ", err)
				}
			default:
				klog.Error("access level error")
				continue
			}
		}

		// 将文件信息存入响应对象
		resp.Files[index] = &file.DownloadInfo{
			FileName: query_file.FileName,
			FilePath: query_file.FilePath,
			FileSize: uint64(query_file.FileSize),
			CheckSum: query_file.CheckSum,
		}
	}

	return
}
