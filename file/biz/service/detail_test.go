package service

import (
	"context"
	"testing"
	file "file/kitex_gen/file"
)

func TestDetail_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDetailService(ctx)
	// init req and assert value

	req := &file.DetailRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
