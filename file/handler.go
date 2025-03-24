package main

import (
	"context"
	file "file/kitex_gen/file"
	"file/biz/service"
)

// FileServiceImpl implements the last service interface defined in the IDL.
type FileServiceImpl struct{}

// Upload implements the FileServiceImpl interface.
func (s *FileServiceImpl) Upload(ctx context.Context, req *file.UploadRequest) (resp *file.UploadResponse, err error) {
	resp, err = service.NewUploadService(ctx).Run(req)

	return resp, err
}

// Download implements the FileServiceImpl interface.
func (s *FileServiceImpl) Download(ctx context.Context, req *file.DownloadRequest) (resp *file.DownloadResponse, err error) {
	resp, err = service.NewDownloadService(ctx).Run(req)

	return resp, err
}

// Check implements the FileServiceImpl interface.
func (s *FileServiceImpl) Check(ctx context.Context, req *file.CheckRequest) (resp *file.CheckResponse, err error) {
	resp, err = service.NewCheckService(ctx).Run(req)

	return resp, err
}

// Delete implements the FileServiceImpl interface.
func (s *FileServiceImpl) Delete(ctx context.Context, req *file.DeleteRequest) (resp *file.DeleteResponse, err error) {
	resp, err = service.NewDeleteService(ctx).Run(req)

	return resp, err
}

// Detail implements the FileServiceImpl interface.
func (s *FileServiceImpl) Detail(ctx context.Context, req *file.DetailRequest) (resp *file.DetailResponse, err error) {
	resp, err = service.NewDetailService(ctx).Run(req)

	return resp, err
}

// Update implements the FileServiceImpl interface.
func (s *FileServiceImpl) Update(ctx context.Context, req *file.UpdateRequest) (resp *file.UpdateResponse, err error) {
	resp, err = service.NewUpdateService(ctx).Run(req)

	return resp, err
}
