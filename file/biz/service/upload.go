package service

import (
	"context"
	"file/biz/dal/model"
	"file/biz/dal/mysql"
	"file/biz/dal/query"
	"file/biz/utils"
	file "file/kitex_gen/file"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
)

type UploadService struct {
	ctx context.Context
} // NewUploadService new UploadService
func NewUploadService(ctx context.Context) *UploadService {
	return &UploadService{ctx: ctx}
}

// Run create note info
func (s *UploadService) Run(req *file.UploadRequest) (resp *file.UploadResponse, err error) {
	// Finish your business logic.

	resp = &file.UploadResponse{
		FileId: make([]string, len(req.Lists)),
	}

	multfile_query := query.Use(mysql.DB)

	for index, upload_list := range req.Lists {
		if upload_list.FileName == "" || upload_list.FilePath == "" || upload_list.OwnerId == "" {
			continue
		}

		uuidv4 := uuid.New().String()

		password_hash := ""
		if upload_list.FilePassword != "" {
			password_hash, err = utils.GenerateHashPassword(upload_list.FilePassword)
			if err != nil {
				klog.Error("bcrypt error: ", err)
				continue
			}
		}

		upload_file, err := os.Open(upload_list.FilePath)
		if err != nil {
			klog.Error("open file error: ", err)
			continue
		}
		defer upload_file.Close()

		file_info, err := upload_file.Stat()
		if err != nil {
			klog.Error("stat file error: ", err)
			continue
		}

		file_size := file_info.Size()
		if file_size == 0 {
			klog.Error("file size is zero")
			continue
		}

		check_sum, err := utils.GenerateFileMD5(upload_file, file_size)
		if err != nil {
			klog.Error("generate file md5 error: ", err)
			continue
		}

		err = multfile_query.File.WithContext(s.ctx).Create(&model.File{
			FileID:           uuidv4,
			FileName:         upload_list.FileName,
			FilePath:         upload_list.FilePath,
			FileSize:         file_size,
			OwnerID:          upload_list.OwnerId,
			AccessLevel:      int32(upload_list.AccessLevel),
			FilePasswordHash: password_hash,
			CheckSum:         check_sum,
		})
		if err != nil {
			klog.Error("create file error: ", err)
			continue
		}

		resp.FileId[index] = uuidv4
	}

	return
}
