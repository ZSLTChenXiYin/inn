package service

import (
	"context"
	"testing"
	file "file/kitex_gen/file"
)

func TestUpload_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUploadService(ctx)
	// init req and assert value

	req := &file.UploadRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
