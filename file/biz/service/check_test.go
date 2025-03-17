package service

import (
	"context"
	"testing"
	file "file/kitex_gen/file"
)

func TestCheck_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckService(ctx)
	// init req and assert value

	req := &file.CheckRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
