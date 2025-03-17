package service

import (
	"context"
	"testing"
	file "file/kitex_gen/file"
)

func TestDelete_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteService(ctx)
	// init req and assert value

	req := &file.DeleteRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
